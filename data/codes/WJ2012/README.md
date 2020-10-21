# 丽泽湖毒素

本题为交互式问题。
交互式问题，选手需要编写程序按照题目要求与评测程序进行互动，并根据互动中获得的信息解出答案。
请在每次输出后强制刷新缓冲区，以实现实时交互。使用C++的选手请使用fflush(stdout)函数，使用Java的选手请使用System.out.flush()方法。
交互式问题仅有Accepted和Wrong Answer两种正常返回，其他所有导致程序出错的情况将返回Runtime Error。

为了办好今年的BNUZCPC开幕式，大仙特地从GDCPC的举办方SCAU那买了一大堆华农酸奶以招待来访的选手们。然而，就在开幕式开始之前，大仙得知，由于广珠城轨动车组列车遭到泛滥的丽泽湖水怪入侵，有一罐华农酸奶在运输过程中被丽泽湖水怪毒素污染。由于丽泽湖水怪毒素剧毒无比，大仙必须在大会开始前找出这一罐被污染的酸奶。
幸运的是，大仙的生物技术同学肥伦有10只可以用于验毒的小白鼠。小白鼠喝到正常酸奶不会有事，但如果喝到有毒素的酸奶，不论多少，都会死亡。为了严谨实验，所有小白鼠不论死活，只能被使用一次。因此大仙需要让不同的小白鼠喝不同罐内的酸奶（方法是将不同罐内的酸奶混合起来给一只小白鼠喝不过这不重要），然后根据哪些小白鼠死亡了，判断哪一罐酸奶被污染了。
为了全体参会人员的安全，请开始你的实验。

## Input
首先，评测程序会输出一行：m cans SCAU yogurt，代表有m罐华农酸奶（1 ≤ m ≤ 1000）。酸奶编号从1到m。
随后用户程序需要输出m行，第mi行表示哪些老鼠喝了第i罐酸奶。每行第一个数字p为喝到这罐酸奶的老鼠的数量，随后p个数字表示喝到这罐酸奶的老鼠的编号（老鼠编号从1到10），数字之间使用空格隔开。
接下来，评测程序会输出两行。第一行评测程序会输出：n mice died，代表有n只老鼠死亡。第二行，评测程序会输出n个数，表示死亡的n只老鼠的编号，数字之间使用空格隔开。
在收到上述信息后，用户程序需要输出一个数，代表被污染的华农酸奶的编号。

## Output
输入样例与输出样例分别为一组的用户程序输出、裁判程序输出，请对照参考。

## simple
```
Your program      | Checker
-----------------------------------------
                  | 3 cans SCAU yogurt
1 1               |           
2 1 2             |
3 1 2 3           |
                  | 3 mice died
                  | 1 2 3
3                 |
------------------------------------------
```