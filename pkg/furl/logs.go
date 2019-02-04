package furl

import (
	"io"
	"k8s.io/api/core/v1"
	"os"
)

func Logs(pod string, container string, follow bool) {

	namespace := client.NamespaceInConfig()

	logOptions := v1.PodLogOptions{
		Container: container,
		Follow:    follow,
	}

	log := client.clientset.CoreV1().Pods(namespace).GetLogs(pod, &logOptions)

	readCloser, err := log.Stream()
	if err != nil {
		panic(err.Error())
	}

	defer readCloser.Close()

	io.Copy(os.Stdout, readCloser)
}
