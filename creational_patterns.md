# Creational Patterns

## Taka Wang

---

## Singleton - Intent

- Ensure a class has **only one instance**, and provide a **global point** of access to it. [GoF] 
- Encapsulated "just-in-time initialization" or "initialization on first use". [[Source Making](https://sourcemaking.com/design_patterns/singleton)]
- Lazy initialization and global point of access. [[Source Making](https://sourcemaking.com/design_patterns/singleton)]

---

## Singleton - Motivation

如何保證一個 Class 只有一個 instance 並且這個 instance 易於被訪問呢？讓 Class 自身負責保存它的**唯一** instance。這個 Class 可以保證沒有其他 instance 被創建，並且它可以提供一個訪問該 instance 的方法。

---

## Singleton - Solution

- Declare the instance as a **private static data member**
- Provide a **public static member function** that encapsulate all initialization code (like a static factory).
- Provides access to the instace.

1. Constructor 為 private，用戶無法透過 new 去 instantiate 它
2. 提供一個自身的靜態私有變量
3. 提供一個公有的靜態工廠方法

---

## Singleton - Check List

1. Define a **private static attribute** in the "single instance" class.
2. Define a **public accessor function** in the class.
3. Do "Lazy initialization" (creation on first use) in the accessor function.
4. Define all constructors to be **protected or private**.
5. Clients may only use the accessor functions to manipulate the Singleton (usually static functions).


---

## Singleton - Naive implementation in "C#"

```csharp
using System;

public class Singleton
{
   private static Singleton instance;

   private Singleton() {}

   public static Singleton Instance
   {
      get
      {
         if (instance == null) // not thread safe
         {
            instance = new Singleton();
         }
         return instance;
      }
   }
}
```
#### Not thread safe!

---

## Double-checked locking

為了減輕 Lock 的效能消耗，先檢查條件**(1)**，有必要再拿鎖**(2)**。但此二步驟之間，可能有其他人已經修改了狀態，所以做第二次的檢查，才開始執行邏輯。
輕量化的鎖，常用在 multithread 下的 lazy initialization。

```c
// pseudo code
if check() { // Lock hint
    lock() {
        if check() {
            // perform your lock-safe code here
        }
    }
}
```

---

## Syntax highlighting

It’s _**sooo easy**_ to show `codeSamples();` in your presentations. Deckset applies syntax highlighting and scales the type size so it always looks great.

---

### Hello World!

```javascript
function myFunction(){
	alert(“Hello World!”)
};
```

### **loooong** lines are scaled down

```objectivec 
UIView *someView = [[UIView alloc] init];
NSString *string = @"some string that is really, really, really really long, and has many, many, many words";
```
