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

#### Example sets

I've successfully used the following set:

  * [Google’s Wikilinks Corpus](https://code.google.com/p/wiki-links/) - [Blog](http://googleresearch.blogspot.se/2013/03/learning-from-big-data-40-million.html)

Example:

    {
      "Digits":[
        {
          "LeadingDigit":0,
          "Count":0,
          "Estimate":0,
          "Dataset":0
        },
        {
          "LeadingDigit":1,
          "Count":3993667,
          "Estimate":30.10299956639812,
          "Dataset":36.573288027760654
        },
        {
          "LeadingDigit":2,
          "Count":1975962,
          "Estimate":17.609125905568124,
          "Dataset":18.09550655022314
        },
        {
          "LeadingDigit":3,
          "Count":1405647,
          "Estimate":12.493873660829994,
          "Dataset":12.872663794041337
        },
        {
          "LeadingDigit":4,
          "Count":840072,
          "Estimate":9.691001300805642,
          "Dataset":7.693229110002648
        },
        {
          "LeadingDigit":5,
          "Count":636960,
          "Estimate":7.918124604762482,
          "Dataset":5.8331657452067045
        },
        {
          "LeadingDigit":6,
          "Count":550225,
          "Estimate":6.694678963061322,
          "Dataset":5.038862129735556
        },
        {
          "LeadingDigit":7,
          "Count":510721,
          "Estimate":5.799194697768673,
          "Dataset":4.677091563925071
        },
        {
          "LeadingDigit":8,
          "Count":492881,
          "Estimate":5.115252244738129,
          "Dataset":4.513716035015112
        },
        {
          "LeadingDigit":9,
          "Count":513480,
          "Estimate":4.575749056067514,
          "Dataset":4.702357992415126
        }
      ]
    }
