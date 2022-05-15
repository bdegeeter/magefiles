package tests_test

import (
	"os"
	"testing"

	"get.porter.sh/magefiles/tests"
	"get.porter.sh/magefiles/tools"
	"github.com/carolynvs/magex/pkg/gopath"
	"github.com/carolynvs/magex/shx"
	"github.com/carolynvs/magex/xplat"
	"github.com/stretchr/testify/require"
)

func TestCreateTestCluster(t *testing.T) {
	tools.EnsureKind()
	clusterName := "test-create-cluster"
	os.Setenv("KIND_NAME", clusterName)
	defer os.Unsetenv("KIND_NAME")
	os.Setenv("KIND_CFG_TEMPLATE", "./testdata/kind.test-cluster.yaml.tmpl")
	defer os.Unsetenv("KIND_CFG_TEMPLATE")
	tests.CreateTestCluster()
	defer tests.DeleteTestCluster()
	xplat.PrependPath(gopath.GetGopathBin())
	err := shx.Run("kind", "get", "kubeconfig", "--name", clusterName)
	require.NoError(t, err)
}
