This is a simple and naïve test to see which implementations of directory traversing is faster. DON'T take this too seriously.

[blog post in Chinese](https://blog.lilydjwg.me/2018/6/14/walk-a-directory-python-vs-go-vs-rust.212932.html)

benchmark.toml is for the tool from [swapview](https://github.com/lilydjwg/swapview).

Result in my case:

About 10k files:
<pre>
<span style="font-weight:bold;color:green;">                    Rust</span>: top: <span style="font-weight:bold;color:gray;">   6.60</span>, min:    6.44, avg:    6.84, max:    7.54, mdev:    0.30, cnt:  20
<span style="font-weight:bold;color:green;">                  Go_3rd</span>: top: <span style="font-weight:bold;color:gray;">   9.44</span>, min:    9.37, avg:    9.57, max:   10.22, mdev:    0.20, cnt:  20
<span style="font-weight:bold;color:green;">                    find</span>: top: <span style="font-weight:bold;color:gray;">  14.52</span>, min:   14.31, avg:   14.73, max:   15.31, mdev:    0.28, cnt:  20
<span style="font-weight:bold;color:green;">                      fd</span>: top: <span style="font-weight:bold;color:gray;">  19.15</span>, min:   16.49, avg:   21.18, max:   26.62, mdev:    2.40, cnt:  20
<span style="font-weight:bold;color:green;">                      Go</span>: top: <span style="font-weight:bold;color:gray;">  26.92</span>, min:   26.77, avg:   27.51, max:   30.12, mdev:    0.91, cnt:  20
<span style="font-weight:bold;color:green;">                  Python</span>: top: <span style="font-weight:bold;color:gray;">  33.64</span>, min:   33.18, avg:   35.19, max:   45.14, mdev:    2.75, cnt:  20
<span style="font-weight:bold;color:green;">                 Python2</span>: top: <span style="font-weight:bold;color:gray;">  38.40</span>, min:   37.56, avg:   39.43, max:   43.84, mdev:    1.66, cnt:  20
<span style="font-weight:bold;color:green;">                 Crystal</span>: top: <span style="font-weight:bold;color:gray;">  47.81</span>, min:   46.92, avg:   51.73, max:   77.20, mdev:    7.91, cnt:  20
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
