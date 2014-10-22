# Working with Data

This chapter outlines techniques for effectively working with data in Python.
This topic is very broad: effective handling of data in Python is less about applying recipes than it is about attaining a basic fluency in the tools available, mainly the NumPy and Pandas packages.
As such, this chapter will be a bit different than the rest of the book: the "recipes" could perhaps better be described as short chapters covering a particular topic.

Throughout, we'll be using the standard abbreviations for common scientific packages, with imports as follows:


```
import numpy as np
import pandas as pd
import scipy as sp
import matplotlib.pyplot as plt
```

As always, more information can be found either in the online documentation, or by using the *Docstrings*, or documentation strings, included in the source code.

Online documentation for these three packages can be found here:

- **NumPy:** http://www.numpy.org
- **Pandas:** http://pandas.pydata.org/
- **SciPy:** http://www.scipy.org
- **Matplotlib:** http://www.matplotlib.org

Often more useful in practice is the docstring, which can be inspected in one of several ways.
See the sidebar (TODO: reference?) for information on finding help.
## Sidebar: Python Docstrings

One of the most important skills for a scientific Python user is knowing how to find help on the various functionality available in Python and its scientific packages.
The standard approach to Python documentation is *docstrings*. These are strings which appear at the beginning of a module, class, or function, and can be examined via several standard methods.
The standard Python convention is for the docstring to be contained within triple-quotes, and they often look as follows:


```
def my_multiply(a, b):
    """Compute the product of a and b
    
    Parameters
    ----------
    a, b: int or float
        numbers to multiply
    """
    return a * b
```

The triple quote denotes a multi-line string literal, and the contents of this string are stored in the function's docstring.  The docstring can be accessed in several ways. At the lowest level, you can access it directly with the ``__doc`` attribute:


```
print(my_multiply.__doc__)
```

More generally, the built-in ``help()`` function can be used to access the docstring along with other useful information about the object:


```
help(my_multiply)
```

Finally, within the IPython terminal or the IPython notebook, a shortcut to the ``help`` function is to add a single question mark after the object:

``` ipython
In [4]: my_multiply?
```

As a result, you will see something like this:
``` python
Type:        function
String form: <function my_multiply at 0x1076f9830>
File:        /Users/jakevdp/Cookbook/01-Working-with-Data/<ipython-input-2-4bec4464217e>
Definition:  my_multiply(a, b)
Docstring:
Compute the product of a and b

Parameters
----------
a, b: int or float
    numbers to multiply
```

This will open a screen with the documentation for the function.
For those who wish to go even deeper, IPython makes it possible to quickly inspect the source code of any object by appending a double question mark:

``` ipython
In [5]: my_multiply??
```
The result is something like this:
``` python
Type:        function
String form: <function my_multiply at 0x1076f9830>
File:        /Users/jakevdp/Cookbook/01-Working-with-Data/<ipython-input-2-4bec4464217e>
Definition:  my_multiply(a, b)
Source:
def my_multiply(a, b):
    """Compute the product of a and b
    
    Parameters
    ----------
    a, b: int or float
        numbers to multiply
    """
    return a * b
```
Be aware that for objects not defined in pure Python (i.e. those directly using the C API), this source-code examination shortcut will only show the docstring.

As you go through the recipes in this book, make good use of these help functionalities!
Most of the tools we will cover have many additional options that we don't have the space to mention or discuss. Use the docstrings to examine these functions and to later remind yourself how to use them!

