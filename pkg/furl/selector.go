package furl

import (
	//"fmt"
	//"gopkg.in/AlecAivazis/survey.v1"
	//"io"
	//"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"gopkg.in/AlecAivazis/survey.v1"
	"k8s.io/api/core/v1"
)

type Selection struct {
	pod       v1.Pod
	container v1.ContainerStatus
}

var client = Client{}

func GetPod(fetchcontainer bool) (v1.Pod, v1.ContainerStatus) {

	client.GetClient()
	namespace := client.NamespaceInConfig()

	pods, err := client.clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	podmap := make(map[string]v1.Pod)
	for _, element := range pods.Items {
		podmap[element.Name] = element
	}

	var podlist []string
	for _, element := range pods.Items {
		podlist = append(podlist, element.Name)
	}

	choice := ""
	prompt := &survey.Select{
		Message: "Choose a pod:",
		Options: podlist,
	}
	survey.AskOne(prompt, &choice, nil)

	var container v1.ContainerStatus
	if fetchcontainer {
		container = getContainer(podmap[choice])
	}

	return podmap[choice], container
}

func getContainer(pod v1.Pod) v1.ContainerStatus {

	// Check for array size

	containermap := make(map[string]v1.ContainerStatus)
	for _, element := range pod.Status.ContainerStatuses {
		containermap[element.Name] = element
	}

	var containerlist []string
	for _, element := range pod.Status.ContainerStatuses {
		containerlist = append(containerlist, element.Name)
	}

	choice := ""
	prompt := &survey.Select{
		Message: "Choose a container:",
		Options: containerlist,
	}
	survey.AskOne(prompt, &choice, nil)

	return containermap[choice]
}
