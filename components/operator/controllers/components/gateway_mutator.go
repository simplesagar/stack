package components

import (
	"context"

	componentsv1beta2 "github.com/formancehq/operator/apis/components/v1beta2"
	apisv1beta2 "github.com/formancehq/operator/pkg/apis/v1beta2"
	"github.com/formancehq/operator/pkg/controllerutils"
	. "github.com/formancehq/operator/pkg/typeutils"
	pkgError "github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

//+kubebuilder:rbac:groups=components.formance.com,resources=gateways,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=components.formance.com,resources=gateways/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=components.formance.com,resources=gateways/finalizers,verbs=update

type GatewayMutator struct {
	Client client.Client
	Scheme *runtime.Scheme
}

func (m *GatewayMutator) SetupWithBuilder(mgr ctrl.Manager, builder *ctrl.Builder) error {
	builder.
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Owns(&networkingv1.Ingress{})
	return nil
}

func (m *GatewayMutator) Mutate(ctx context.Context, gateway *componentsv1beta2.Gateway) (*ctrl.Result, error) {
	apisv1beta2.SetProgressing(gateway)

	_, err := m.reconcileDeployment(ctx, gateway)
	if err != nil {
		return controllerutils.Requeue(), pkgError.Wrap(err, "Reconciling deployment")
	}

	apisv1beta2.SetReady(gateway)

	return nil, nil
}

func (m *GatewayMutator) reconcileDeployment(ctx context.Context, gateway *componentsv1beta2.Gateway) (*appsv1.Deployment, error) {
	matchLabels := CreateMap("app.kubernetes.io/name", "gateway")

	var env []corev1.EnvVar

	if gateway.Spec.Monitoring != nil {
		env = append(env, gateway.Spec.Monitoring.Env("")...)
	}

	ret, _, err := controllerutils.CreateOrUpdate(ctx, m.Client, client.ObjectKeyFromObject(gateway),
		controllerutils.WithController[*appsv1.Deployment](gateway, m.Scheme),
		func(deployment *appsv1.Deployment) error {
			deployment.Spec = appsv1.DeploymentSpec{
				Selector: &metav1.LabelSelector{
					MatchLabels: matchLabels,
				},
				Template: corev1.PodTemplateSpec{
					ObjectMeta: metav1.ObjectMeta{
						Labels: matchLabels,
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{{
							Name:            "gateway",
							Image:           controllerutils.GetImage("gateway", gateway.Spec.Version),
							ImagePullPolicy: controllerutils.ImagePullPolicy(gateway.Spec),
							Env:             env,
							Ports: []corev1.ContainerPort{{
								Name:          "http",
								ContainerPort: 80,
							}},
						}},
					},
				},
			}
			return nil
		})
	if err != nil {
		return nil, err
	}

	return ret, err
}

var _ controllerutils.Mutator[*componentsv1beta2.Gateway] = &GatewayMutator{}

func NewGatewayMutator(client client.Client, scheme *runtime.Scheme) controllerutils.Mutator[*componentsv1beta2.Gateway] {
	return &GatewayMutator{
		Client: client,
		Scheme: scheme,
	}
}
