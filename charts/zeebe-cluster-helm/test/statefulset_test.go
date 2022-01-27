package test

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	appsv1 "k8s.io/api/apps/v1"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
)

type statefulSetTemplateTest struct {
	suite.Suite
	chartPath string
	release   string
	namespace string
	templates []string
}

func TestStatefulSetTemplate(t *testing.T) {
	t.Parallel()

	chartPath, err := filepath.Abs("../")
	require.NoError(t, err)

	suite.Run(t, &statefulSetTemplateTest{
		chartPath: chartPath,
		release:   "zeebe-cluster-helm",
		namespace: "zeebe-" + strings.ToLower(random.UniqueId()),
		templates: []string{"templates/statefulset.yaml"},
	})
}

func (s *statefulSetTemplateTest) TestContainerSpecImage() {
	options := &helm.Options{
		SetValues: map[string]string{
			"image.repository": "helm/zeebe",
			"image.tag":        "a.b.c",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", s.namespace),
	}
	output := helm.RenderTemplate(s.T(), options, s.chartPath, s.release, s.templates)

	var statefulSet appsv1.StatefulSet
	helm.UnmarshalK8SYaml(s.T(), output, &statefulSet)

	expectedContainerImage := "helm/zeebe:a.b.c"
	containers := statefulSet.Spec.Template.Spec.Containers
	s.Require().Equal(len(containers), 1)
	s.Require().Equal(containers[0].Image, expectedContainerImage)
}
