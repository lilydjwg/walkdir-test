This is a simple and naïve test to see which implementations of directory traversing is faster. DON'T take this too seriously.

[blog post in Chinese](https://blog.lilydjwg.me/2018/6/14/walk-a-directory-python-vs-go-vs-rust.212932.html)

benchmark.toml is for the tool from [swapview](https://github.com/lilydjwg/swapview).

Result in my case:

About 10k files:
<pre>
<span style="color:green;font-weight:bold;">   Rust</span>: top: <span style="color:gray;font-weight:bold;">   4.78</span>, min:    4.72, avg:    4.90, max:    5.46, mdev:    0.17, cnt:  20
<span style="color:green;font-weight:bold;"> Go_3rd</span>: top: <span style="color:gray;font-weight:bold;">   7.71</span>, min:    7.64, avg:    7.79, max:    8.41, mdev:    0.16, cnt:  20
<span style="color:green;font-weight:bold;">   find</span>: top: <span style="color:gray;font-weight:bold;">  11.49</span>, min:   11.32, avg:   11.76, max:   14.18, mdev:    0.59, cnt:  20
<span style="color:green;font-weight:bold;">     fd</span>: top: <span style="color:gray;font-weight:bold;">  18.17</span>, min:   15.18, avg:   21.29, max:   29.94, mdev:    3.84, cnt:  20
<span style="color:green;font-weight:bold;">     Go</span>: top: <span style="color:gray;font-weight:bold;">  21.08</span>, min:   20.91, avg:   21.28, max:   22.70, mdev:    0.37, cnt:  20
<span style="color:green;font-weight:bold;"> Python</span>: top: <span style="color:gray;font-weight:bold;">  29.66</span>, min:   29.51, avg:   30.43, max:   35.84, mdev:    1.45, cnt:  20
<span style="color:green;font-weight:bold;">Python2</span>: top: <span style="color:gray;font-weight:bold;">  30.37</span>, min:   30.10, avg:   30.85, max:   33.15, mdev:    0.75, cnt:  20
</pre>

About 300k files:
<pre>
<span style="color:green;font-weight:bold;">     fd</span>: top: <span style="color:gray;font-weight:bold;"> 265.80</span>, min:  259.84, avg:  273.89, max:  319.76, mdev:   15.03, cnt:  20
<span style="color:green;font-weight:bold;">   Rust</span>: top: <span style="color:gray;font-weight:bold;"> 269.98</span>, min:  266.86, avg:  272.82, max:  282.84, mdev:    4.17, cnt:  20
<span style="color:green;font-weight:bold;"> Go_3rd</span>: top: <span style="color:gray;font-weight:bold;"> 361.17</span>, min:  359.05, avg:  363.82, max:  370.22, mdev:    3.31, cnt:  20
<span style="color:green;font-weight:bold;">   find</span>: top: <span style="color:gray;font-weight:bold;"> 454.03</span>, min:  450.79, avg:  458.51, max:  467.31, mdev:    5.08, cnt:  20
<span style="color:green;font-weight:bold;"> Python</span>: top: <span style="color:gray;font-weight:bold;"> 624.80</span>, min:  615.73, avg:  630.67, max:  640.88, mdev:    6.79, cnt:  20
<span style="color:green;font-weight:bold;">     Go</span>: top: <span style="color:gray;font-weight:bold;"> 890.03</span>, min:  876.98, avg:  910.63, max:  967.14, mdev:   24.84, cnt:  20
<span style="color:green;font-weight:bold;">Python2</span>: top: <span style="color:gray;font-weight:bold;">1171.38</span>, min: 1157.19, avg: 1189.99, max: 1228.09, mdev: 4186.28, cnt:  20
</pre>
