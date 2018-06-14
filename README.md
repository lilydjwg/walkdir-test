This is a simple and naïve test to see which implementations of directory traversing is faster. DON'T take this too seriously.

[blog post in Chinese](https://blog.lilydjwg.me/2018/6/14/walk-a-directory-python-vs-go-vs-rust.212932.html)

benchmark.toml is for the tool from [swapview](https://github.com/lilydjwg/swapview).

Result in my case:

<pre>
<span style="color:green;font-weight:bold;">                    Rust</span>: top: <span style="color:gray;font-weight:bold;">   4.78</span>, min:    4.72, avg:    4.90, max:    5.46, mdev:    0.17, cnt:  20
<span style="color:green;font-weight:bold;">                  Go_3rd</span>: top: <span style="color:gray;font-weight:bold;">   7.71</span>, min:    7.64, avg:    7.79, max:    8.41, mdev:    0.16, cnt:  20
<span style="color:green;font-weight:bold;">                    find</span>: top: <span style="color:gray;font-weight:bold;">  11.49</span>, min:   11.32, avg:   11.76, max:   14.18, mdev:    0.59, cnt:  20
<span style="color:green;font-weight:bold;">                      fd</span>: top: <span style="color:gray;font-weight:bold;">  18.17</span>, min:   15.18, avg:   21.29, max:   29.94, mdev:    3.84, cnt:  20
<span style="color:green;font-weight:bold;">                      Go</span>: top: <span style="color:gray;font-weight:bold;">  21.08</span>, min:   20.91, avg:   21.28, max:   22.70, mdev:    0.37, cnt:  20
<span style="color:green;font-weight:bold;">                  Python</span>: top: <span style="color:gray;font-weight:bold;">  29.66</span>, min:   29.51, avg:   30.43, max:   35.84, mdev:    1.45, cnt:  20
<span style="color:green;font-weight:bold;">                 Python2</span>: top: <span style="color:gray;font-weight:bold;">  30.37</span>, min:   30.10, avg:   30.85, max:   33.15, mdev:    0.75, cnt:  20
</pre>
