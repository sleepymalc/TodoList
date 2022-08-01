# TODO

<p align="center"><b><i>
	A Small TODO list Project
</i></b></p>

## Abstract

This project is based on the following elements: Golang, Gin Framework, JWT (JSON Web Token), Restful API, MongoDB, GCP (Google Cloud Platform), GKE (Google Kubernetes Engine). We use Postman to test all the API functionality. Additionally, one can build with Docker and push it to GKE and expose it to the internet.

## Features

1. users/signup: Post method
  
    Sample Input

    ```JSON
    {
      "user_id": "admin",
      "password": "88888888"
    }
    ```

    Sample Output

    ```JSON
    {
      "InsertedID": "60581eb9cc2216b508d7477d"
    }
    ```

2. users/login : Post method
  
    Sample Input

    ```JSON
    {
      "user_id": "admin",
      "password": "88888888"
    }
    ```

    Sample Output

    ```JSON
    {
      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyX2lkIjoiYWRtaW4iLCJleHAiOjE2MTY0NzQxNjl9.aAlh_LuxqLfWWSzCd3uA3C2RoBOTnC3HSqPaAvzYIkE"
    }
    ```

3. users/todo_list: Post method

    Sample Input

    ```JSON
    {
      "todo_list": ["do homework", "write diary"]
    }
    ```

    Sample Output

    ```JSON
    {
      "MatchedCount": 1,
      "ModifiedCount": 1,
      "UpsertedCount": 0,
      "UpsertedID": null
    }
    ```

4. user/todo_list: Get method

    Sample Output

    ```JSON
    {
      "do homework",
      "write diary"
    }
    ```

    > Note that this method does not have input.

5. user/todo_list: Delete method

    Sample Input

    ```JSON
    {
      "delete_element": "do homework"
    }
    ```

    Sample Output

    ```JSON
    {
      "MatchedCount": 1,
      "ModifiedCount": 1,
      "UpsertedCount": 0,
      "UpsertedID": null
    }
    ```

## GCP + GKE

To build a docker file, first, go to `/Todo_List`, and then type in the following command:

```bash
> docker build -t asia.gcr.io/PROJECT_ID/todo_list .
```

which will generate an image of this project. To push the image to GCP, we type in the following command:

```bash
> docker push asia.gcr.io/PROJECT_ID/todo_list
```

which just pushes the image file to your GCP project. Now, to run the project o GKE, we use:

```bash
> kubectl run todolist --image=asia.gcr.io/PROJECT_ID/todo_list
```

which tells GKE to run the project. You can check if this is working by the following command:

```bash
> kubectl get pods
```

this will give you the current state of your `pod`. If the state is running, now you can enter the Docker by the following command:

```bash
> kubectl exec -it todolist -- bash
```

which will let you get into the docker. Finally, to expose your project to the Internet, use the following command:

```bash
> kubectl expose deployment todolist-server --type LoadBalancer --port 80 --target-port 80
```

which specifies both the exposed port in docker and the port of your localhost is 80, and also generates a load-balancer. Now, use the following command to truly deploy the project:

```bash
> kubectl create deployment todolist-server --image=asia.gcr.io/PROJECT_ID/todo_list
```

To get the IP of your project, use the following command:

```bash
> kubectl get service
```

To delete the service, use:

```bash
> tubectl delete service todolist-service
```

## Postman

In the repo, you can find `TodoList.postman_collection.JSON` and then import to the Postman.
