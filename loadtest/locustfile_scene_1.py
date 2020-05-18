from locust import HttpLocust, TaskSet, task, between


class UserABehavior(TaskSet):
    def on_start(self):
        self.warmup()

    def warmup(self):
        self.client.get("/")

    @task(1)
    def home(self):
        self.client.get("/")

    @task(2)
    def volatile(self):
        self.client.get("/volatile")


class ApiUser(HttpLocust):
    task_set = UserABehavior
    wait_time = between(5, 9)
