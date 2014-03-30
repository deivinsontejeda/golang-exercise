This code has as proposal to help to undestand better the pointer in Go.

### Output 
```
1. main  -- i  int: &i=0x2101f1018 i=42
2. main  -- p *int: &p=0x2101f1020 p=&i=0x2101f1018  *p=i=42
3. byval -- q *int: &q=0x2101f1028 q=&i=0x2101f1018  *q=i=42
4. byval -- q *int: &q=0x2101f1028 q=&i=0x2101f1018  *q=i=4143
5. main  -- p *int: &p=0x2101f1020 p=&i=0x2101f1018  *p=i=4143
6. main  -- i  int: &i=0x2101f1018 i=4143
```

### What happened?
In function ``main``, ``i`` is an ``int`` variable at memory location (``&i``) ``0x2101f1018`` with an initial value ``42``.

In function ``main``, ``p`` is a pointer to an ``int`` variable at memory location (``&p``) ``0x2101f1020`` with a value (``p=&i``) ``0x2101f1018`` which points to an ``int`` value (``*p=i``) ``42``.

In function ``main``, ``byval(p)`` is a function call which assigns the value (``p=&i``) ``0x2101f1018`` of the argument at memory location (``&p``) ``0x2101f1020`` to the function ``byval`` parameter ``q`` at memory location (&q) ``0x2101f1028``. In other words, memory is allocated for the ``byval`` parameter ``q`` and the value of the main byval argument ``p`` is assigned to it; the values of ``p`` and ``q`` are initially the same, but the variables ``p`` and ``q`` are distinct.

In function ``byval``, using pointer ``q`` (``*int``), which is a copy of pointer ``p`` (``*int``), integer ``*q`` (``i``) is set to a new int value ``4143``. At the end before returning. the pointer ``q`` is set to nil (zero value), which has no effect on ``p`` since ``q`` is a copy.

In function ``main``, ``p`` is a pointer to an int variable at memory location (``&p``) ``0x2101f1020`` with a value (``p=&i``) ``0x2101f1018`` which points to a new int value (``*p=i``) ``4143``.

In function main, ``i`` is an int variable at memory location (``&i``) ``0x2101f1018`` with a final value (``i``) ``4143``.
