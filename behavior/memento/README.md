## 备忘录模式允许在不暴露对象实现细节的情况下保存和恢复对象之前的状态。
## 原发器（originator) 可以生成自身状态的快照，也可以在需要时通过快照恢复自身状态
## 备忘录（Memento）是原发器状态快照的值对象，通常将备忘录设置为不可变的，并通过构造函数一次性传递数据
## 负责人（Caretaker)仅知道 什么时候，什么原因捕捉原发器状态，以及什么时候恢复状态。负责人通过保存备忘录栈来记录原发器的历史状态。

## 当需要创建对象状态快照来恢复之前的状态时候，可以使用备忘录模式
## 当直接访问对象的成员变量，get或set将导致封装被突破时，可以使用备忘录模式

