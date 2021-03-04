from Base import PriorityQueueBase

class UnsortedPriorityQueue(PriorityQueueBase):

    def _find_min(self):
        # 返回有最小键值的项
        if self.is_empty():
            raise Empty('Priority queue is empty')
        small = self._data.first()
        walk = self._data.after(small)
        while walk is not None:
            if walk.element() < small.element():
                small = walk
            walk = self._data.after(walk)
        return small

    def __init__(self):
        self._data = PositionalList()