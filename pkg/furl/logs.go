package furl

import (
	"io"
	"k8s.io/api/core/v1"
	"os"
)

func Logs(selection Selection, follow bool) {

	namespace := client.GetNamespace()

	tailLines := int64(10)
	logOptions := v1.PodLogOptions{
		Container: selection.Container.Name,
		Follow:    follow,
		TailLines: &tailLines,
	}

	log := client.clientset.CoreV1().Pods(namespace).GetLogs(selection.Pod.Name, &logOptions)

	readCloser, err := log.Stream()
	if err != nil {
		panic(err.Error())
	}

	defer readCloser.Close()
	io.Copy(os.Stdout, readCloser)
}
