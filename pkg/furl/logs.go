package furl

import (
	"io"
	"k8s.io/api/core/v1"
	"os"
)

func Logs(selection Selection, follow bool) {

	namespace := client.GetNamespace()

	logOptions := v1.PodLogOptions{
		Container: selection.container.Name,
		Follow:    follow,
	}

	log := client.clientset.CoreV1().Pods(namespace).GetLogs(selection.pod.Name, &logOptions)

	readCloser, err := log.Stream()
	if err != nil {
		panic(err.Error())
	}

	defer readCloser.Close()

	io.Copy(os.Stdout, readCloser)
}
