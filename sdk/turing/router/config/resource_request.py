from dataclasses import dataclass, field
from typing import Optional

import turing.generated.models
from turing.generated.model_utils import OpenApiModel


@dataclass
class ResourceRequest:
    min_replica: int
    max_replica: int
    cpu_request: str
    memory_request: str
    cpu_limit: Optional[str] = None

    _min_replica: int = field(init=False, repr=False)
    _max_replica: int = field(init=False, repr=False)
    _cpu_request: str = field(init=False, repr=False)
    _memory_request: str = field(init=False, repr=False)
    _cpu_limit: Optional[str] = field(init=False, repr=False, default=None)

    @property
    def min_replica(self) -> int:
        return self._min_replica

    @min_replica.setter
    def min_replica(self, min_replica: int):
        if hasattr(self, "max_replica"):
            ResourceRequest._verify_min_max_replica(min_replica, self.max_replica)
        self._min_replica = min_replica

    @property
    def max_replica(self) -> int:
        return self._max_replica

    @max_replica.setter
    def max_replica(self, max_replica: int):
        if hasattr(self, "min_replica"):
            ResourceRequest._verify_min_max_replica(self.min_replica, max_replica)
        self._max_replica = max_replica

    @property
    def cpu_request(self) -> str:
        return self._cpu_request

    @cpu_request.setter
    def cpu_request(self, cpu_request: str):
        self._cpu_request = cpu_request

    @property
    def cpu_limit(self) -> str:
        return self._cpu_limit

    @cpu_limit.setter
    def cpu_limit(self, cpu_limit):
        if type(cpu_limit) is property:
            cpu_limit = ResourceRequest._cpu_limit
        self._cpu_limit = cpu_limit

    @property
    def memory_request(self) -> str:
        return self._memory_request

    @memory_request.setter
    def memory_request(self, memory_request: str):
        self._memory_request = memory_request

    @classmethod
    def _verify_min_max_replica(cls, min_replica: int, max_replica: int):
        if min_replica > max_replica:
            raise InvalidReplicaCountException(
                f"Min replica must be <= max_replica; "
                f"min_replica passed: {min_replica}, max_replica passed: {max_replica}"
            )

    def to_open_api(self) -> OpenApiModel:
        return turing.generated.models.ResourceRequest(
            min_replica=self.min_replica,
            max_replica=self.max_replica,
            cpu_request=self.cpu_request,
            cpu_limit=self.cpu_limit,
            memory_request=self.memory_request,
        )


class InvalidReplicaCountException(Exception):
    pass
