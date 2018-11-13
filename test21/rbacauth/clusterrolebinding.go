package rbacauth

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/client-go/kubernetes"
	authv1 "k8s.io/api/rbac/v1"
	rbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

type RbacClusterRoleBinding struct {
	Rolebinding rbacv1.ClusterRoleBindingInterface
}
func (self RbacClusterRoleBinding)ClusterRolebindingCreateHandle()rbacv1.ClusterRoleBindingInterface{
	authRolbindings := clientset.RbacV1().ClusterRoleBindings()
	//self.Rolebinding = authRolbindings
	fmt.Println("------------0-------Rolebinding",authRolbindings)
	return authRolbindings
}

func (self RbacClusterRoleBinding) CreateClusterRoleBinding() {
	rolbinding := &authv1.ClusterRoleBinding{

		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{authv1.AutoUpdateAnnotationKey: "true"},
			Labels:      map[string]string{"kubernetes.io/bootstrapping": "rbac-defaults"},
			Name:        "cluster-dl",
			//ResourceVersion:authv1.ResourceAll,//"76",
		},
		RoleRef: authv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "cluster-dl",
		},
		Subjects: []authv1.Subject{
			{
				APIGroup: "rbac.authorization.k8s.io",
				Kind:     authv1.GroupKind,
				Name:     "admin-dl",
			},
		},
	}
	fmt.Println("------------1-------Rolebinding",self.Rolebinding)

	ree, err := self.Rolebinding.Create(rolbinding)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created authClusterRolbindings %q.\n", ree.GetObjectMeta().GetName())
}

func (self RbacClusterRoleBinding) DeleteClusterRoleBindings(rolename string) {
	err:=self.Rolebinding.Delete(rolename,&metav1.DeleteOptions{})

	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted authClusterRolbindings .\n")
}
