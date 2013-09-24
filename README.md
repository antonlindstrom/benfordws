# Benford WS

A web API for Benfordslaw.

Remember that the data must be integers in an array, the array must have the key `Set`:

The following JSON is valid:

    {
      "Set": []
    }

## Examples

Running the set from CLI:: 

    curl -X POST -d "{ \"Set\": [ 1, 5, 8, 100, 43, 300 ] }" http://localhost:6000

Run the Fibonacci numbers in the fibonacci.json:

    curl -X POST -d @fibonacci.json http://localhost:6000

You'll get a JSON hash back with information about the Benford estimate of the numbers.

This will log:

    2013/09/15 15:20:52 Processed dataset of 6 numbers
