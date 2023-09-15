package main

import (
	"github.com/nailabx/velero-image-update/internal/plugin"
	"github.com/sirupsen/logrus"
	"github.com/vmware-tanzu/velero/pkg/plugin/framework"
)

func main() {
	framework.NewServer().
		RegisterRestoreItemAction("nailabx.io/velero-image-updte", newRestorePlugin).
		RegisterRestoreItemAction("nailabx.io/restore-plugin", newRestorePlugin).
		RegisterRestoreItemActionV2("nailabx.io/restore-pluginv2", newRestorePluginV2).
		Serve()
}

func newRestorePlugin(logger logrus.FieldLogger) (interface{}, error) {
	return plugin.NewRestorePlugin(logger), nil
}

func newRestorePluginV2(logger logrus.FieldLogger) (interface{}, error) {
	return plugin.NewRestorePluginV2(logger), nil
}
