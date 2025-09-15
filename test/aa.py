from datetime import datetime
from threading import RLock


class TaskDataSlave:
    def __init__(self, id, name):
        self.id = id
        self.name = name
        self.create_time = datetime.now()
        self.modify_time = datetime.now()


class TaskData:
    def __init__(self, id, name):
        self.id = id
        self.name = name
        self.slave = []
        self.create_time = datetime.now()
        self.modify_time = datetime.now()


class Task:
    def __init__(self):
        self._auto_id = 0
        self._lock = RLock()
        self.task_new = []
        self.task_now = []
        self.task_old = []

    def _next_id(self):
        self._auto_id += 1
        return self._auto_id

    def _with_list(self, list_name):
        return getattr(self, list_name)

    def add_task(self, list_name, name):
        self._lock.acquire()
        try:
            task = TaskData(self._next_id(), name)
            self._with_list(list_name).append(task)
        finally:
            self._lock.release()

    def remove_task(self, list_name, task_id):
        self._lock.acquire()
        try:
            lst = self._with_list(list_name)
            before = len(lst)
            new_list = []
            for t in lst:
                if t.id != task_id:
                    new_list.append(t)
            lst[:] = new_list
            return before - len(lst)
        finally:
            self._lock.release()

    def update_task(self, list_name, task_id, name):
        self._lock.acquire()
        try:
            updated = 0
            for t in self._with_list(list_name):
                if t.id == task_id:
                    t.name = name
                    t.modify_time = datetime.now()
                    updated += 1
            return updated
        finally:
            self._lock.release()

    def get_by_time(self, list_name, start, end, use_create=True):
        self._lock.acquire()
        try:
            result = []
            for t in self._with_list(list_name):
                if use_create:
                    time_field = t.create_time
                else:
                    time_field = t.modify_time
                if start <= time_field <= end:
                    result.append(t)
            return result
        finally:
            self._lock.release()

    def add_slave(self, master_id, name):
        self._lock.acquire()
        try:
            slave = TaskDataSlave(self._next_id(), name)
            added_count = 0
            for lst in (self.task_new, self.task_now, self.task_old):
                for task in lst:
                    if task.id == master_id:
                        task.slave.append(slave)
                        task.modify_time = datetime.now()
                        added_count += 1
            return added_count
        finally:
            self._lock.release()

    def remove_slave(self, master_id, slave_id):
        self._lock.acquire()
        try:
            removed = 0
            for lst in (self.task_new, self.task_now, self.task_old):
                for task in lst:
                    if task.id == master_id:
                        before = len(task.slave)
                        new_slaves = []
                        for s in task.slave:
                            if s.id != slave_id:
                                new_slaves.append(s)
                        task.slave[:] = new_slaves
                        removed += before - len(task.slave)
                        task.modify_time = datetime.now()
            return removed
        finally:
            self._lock.release()

    def update_slave(self, master_id, slave_id, name):
        self._lock.acquire()
        try:
            updated = 0
            for lst in (self.task_new, self.task_now, self.task_old):
                for task in lst:
                    if task.id == master_id:
                        for s in task.slave:
                            if s.id == slave_id:
                                s.name = name
                                s.modify_time = datetime.now()
                                updated += 1
                        task.modify_time = datetime.now()
            return updated
        finally:
            self._lock.release()

    def sort_by_time(self, items, use_create=True, reverse=False):
        def key_func_create(t):
            return t.create_time

        def key_func_modify(t):
            return t.modify_time

        if use_create:
            key_func = key_func_create
        else:
            key_func = key_func_modify

        return sorted(items, key=key_func, reverse=reverse)


if __name__ == "__main__":
    import time
    a = time.time()
    t = Task()
    t.add_task("task_new", "学习 Python")
    t.add_slave(1, "啊啊啊")
    t.update_slave(1, 2,"空的反馈")
    t.remove_slave(1, 3)
    tasks_sorted = t.sort_by_time(t.task_new, use_create=True, reverse=True)
    i = 0
    while i < 1000000:
        t.update_slave(1, 2, 'dkkfd' + str(i))
        i += 1

    for task in tasks_sorted:
        print(task.__dict__["slave"])

    print(time.time() - a)