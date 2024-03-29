Vending machines have an array of buckets where each bucket contains a number of (possibly different products). When "vending" a particular product from a particular bucket, only the front-most product can be vended. When a user generates an order containing a variety of different products, the system needs to run an algorithm to decide whether the products can be sold and what the buckets look like after they have been sold.
We encode a single bucket by a string as follows: <int_1>,<int_2>,...,<int_n> where the front-most <int_1> has to be vended first. So e.g. 1,2,1,3,4 would be a bucket in which product 1 can be sold first, then product 2, then again product 1 and so on.
We encode the whole bucket array by concatenating several bucket encodings with ; - so e.g. 1,2,1,3,4;5,2,3,3,3 would be a two-bucket configuration where product 1 can be vended immediately from the first bucket and product 5 can be vended immediately from the second bucket.
A user order is encoded by <int_1>,<int_2>,...,<int_m> where the order of product ids does not matter, e.g. 5,2,2 is the same as 2,2,5.
The algorithm now receives a bucket array as well as a user order as an input and either outputs "IMPOSSIBLE" or outputs a new bucket array after the order has been vended.
Please test your program with the following input:
Buckets 1->100,2->50,3->150,5->200,5->200;2->50,5->200,4->100,3->150,4->100;3->150,5->200,4->100,1->100,1->100;5->200,1->100,1->100,1->100,1->100
Where 1 - name product, 100 - price
Order 1,2,3,4,5