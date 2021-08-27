## 中介者模式减少对象之间混乱无序的依赖关系。该模式会限制对象之间的交互，迫使它们通过一个中介者对象进行合作。

## 当一些对象和其他对象紧密耦合以致难以对其进行修改时，可以使用中介者模式
## 当组件因过于依赖其他组件而无法在不同应用中复用时，使用中介者模式
## 如果为了能在不同场景下复用一些基本行为，导致被迫创建大量组件子类时，可以使用中介者模式

##实现
## 找到一组当前紧密耦合，且提供其独立性能带来更大好处的类
## 声明中介者接口并描述中介者和各种组件之间所需的交流接口
## 实现中介者类 、 中介者可以负责组件的创建，销毁
## 组件必须保存对中介者对象的引用
## 使组件调用中介者的通知方法取代之前调用其他组件的方法。调用其他组建的方法抽离封装到中介者中。