package furl

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"gopkg.in/AlecAivazis/survey.v1"
	"k8s.io/api/core/v1"
)

type Selection struct {
	Pod        v1.Pod
	Container  v1.ContainerStatus
}

var client = Client{}

func GetPod(fetchcontainer bool) Selection {

	client.GetClient()
	namespace := client.GetNamespace()

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
		Message: "Choose a Pod:",
		Options: podlist,
	}
	survey.AskOne(prompt, &choice, nil)

	var container v1.ContainerStatus
	if fetchcontainer {
		container = getContainer(podmap[choice])
	}

	selection := Selection{
		Pod:       podmap[choice],
		Container: container,
	}

	return selection
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
		Message: "Choose a Container:",
		Options: containerlist,
	}
	survey.AskOne(prompt, &choice, nil)

	return containermap[choice]
}
