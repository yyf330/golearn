package rbacauth

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	authv1 "k8s.io/api/rbac/v1"
	rbacv1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

type RbacClusterRole struct {
	ClusterRole rbacv1.ClusterRoleInterface
}
func (self RbacClusterRole)ClusterRoleCreateHandle() rbacv1.ClusterRoleInterface{
	authRole := clientset.RbacV1().ClusterRoles()
	self.ClusterRole = authRole
	return authRole
}

func (self RbacClusterRole) CreateClusterRole() {
	//authClusterRoles := clientset.RbacV1().ClusterRoles()
	clusterRole := &authv1.ClusterRole{
		ObjectMeta: metav1.ObjectMeta{
			Annotations:map[string]string{authv1.AutoUpdateAnnotationKey:"true"},
			Labels:map[string]string{"kubernetes.io/bootstrapping": "rbac-defaults"},
			Name:"cluster-dl",
			//ResourceVersion:"32",
		},
		Rules:[]authv1.PolicyRule{
			{
				APIGroups:[]string{authv1.APIGroupAll},
				Resources:[]string{authv1.ResourceAll},
				Verbs:[]string{authv1.VerbAll},
				//NonResourceURLs:[]string{authv1.NonResourceAll},
			},
		},

	}
	re,err:=self.ClusterRole.Create(clusterRole)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Created authClusterRoles %q.\n", re.GetObjectMeta().GetName())
}

func (self RbacClusterRole) DeleteClusterRole(rolename string) {
	err:=self.ClusterRole.Delete(rolename,&metav1.DeleteOptions{})

	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted authClusterRoles .\n")
}
