from locust import HttpUser, TaskSet, task, between


class UserBBehavior(TaskSet):
    def on_start(self):
        self.warmup()

    def warmup(self):
        self.client.get("/")

    @task(1)
    def volatile(self):
        self.client.get("/buggy")


class AdminUser(HttpUser):
    tasks = {UserBBehavior}
    wait_time = between(5, 15)
