const express = require("express");
const app = express();
const port = 8070;

function fib(n) {
  if (n <= 2) {
    return 1;
  }
  return fib(n - 1) + fib(n - 2);
}

app.get("/fib", (req, res) => {
  const str = req.query.n;

  if (!str) {
    return res.status(400).send("Query parameter 'n' is required");
  }

  const n = parseInt(str, 10);

  if (isNaN(n) || n < 1) {
    return res
      .status(400)
      .send("Query parameter 'n' has to be a positive integer");
  }

  res.set("Content-Type", "text/plain");
  res.send(`The Fibonacci number at position ${n} is ${fib(n)}\n`);
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}...`);
});
