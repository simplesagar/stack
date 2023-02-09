package v1beta2

import (
	authcomponentsv1beta2 "github.com/formancehq/operator/apis/auth.components/v1beta2"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// +kubebuilder:object:generate=true
type GatewaySpec struct {
	// +optional
	Scaling ScalingSpec `json:"scaling,omitempty"`
}

func (in GatewaySpec) NeedAuthMiddleware() bool {
	return false
}

func (in GatewaySpec) Spec(stack *Stack, configuration ConfigurationSpec) any {
	return nil
}

func (in GatewaySpec) HTTPPort() int {
	return 80
}

func (in GatewaySpec) AuthClientConfiguration(stack *Stack) *authcomponentsv1beta2.ClientConfiguration {
	return nil
}

func (in GatewaySpec) Validate() field.ErrorList {
	return field.ErrorList{}
}
