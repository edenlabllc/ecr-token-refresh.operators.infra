package secret

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	ecrv1alpha1 "ecr-token-refresh.operators.infra/api/v1alpha1"
	"ecr-token-refresh.operators.infra/internal/aws"
)

var secretMeta = metav1.TypeMeta{
	APIVersion: "v1",
	Kind:       "Secret",
}

type InputFromCRD struct {
	CRD *ecrv1alpha1.ECRTokenRefresh
}

type Creator interface {
	CreateSecret(input *InputFromCRD) (*v1.Secret, error)
}

type DefaultSecretCreator struct {
	r aws.TokenRetriever
}

type DockerAuths struct {
	Auths map[string]Auth `json:"auths"`
}

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Auth     string `json:"auth"`
}

func NewDefaultSecretCreator(r aws.TokenRetriever) *DefaultSecretCreator {
	return &DefaultSecretCreator{
		r: r,
	}
}

func getDockerAuth(registry, username, password string) ([]byte, error) {
	return json.Marshal(&DockerAuths{
		Auths: map[string]Auth{
			registry: {
				Username: username,
				Password: password,
				Auth:     base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password))),
			},
		},
	})
}

func (sg *DefaultSecretCreator) CreateSecret(input *InputFromCRD) (*v1.Secret, error) {
	var (
		token      string
		err        error
		dockerAuth []byte
	)

	if token, err = sg.r.GetToken(input.CRD.Spec.Region); err != nil {
		return nil, err
	}

	if dockerAuth, err = getDockerAuth(input.CRD.Spec.ECRRegistry, "AWS", token); err != nil {
		return nil, err
	}

	return &v1.Secret{
		TypeMeta: secretMeta,
		ObjectMeta: metav1.ObjectMeta{
			Namespace: input.CRD.ObjectMeta.Namespace,
			Name:      input.CRD.ObjectMeta.Name,
		},
		Data: map[string][]byte{
			".dockerconfigjson": dockerAuth,
		},
		Type: "kubernetes.io/dockerconfigjson",
	}, nil
}
