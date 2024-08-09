from flask import Flask, request, abort

app = Flask(__name__)

def fib(n):
    if n <= 2:
        return 1
    return fib(n - 1) + fib(n - 2)

@app.route('/fib', methods=['GET'])
def handle_fib():
    n_str = request.args.get('n')
    if n_str is None:
        abort(400, description="Query parameter 'n' is required")

    try:
        n = int(n_str)
        if n < 1:
            raise ValueError
    except ValueError:
        abort(400, description="Query parameter 'n' has to be a positive integer")

    return f"The Fibonacci number at position {n} is {fib(n)}\n", 200

if __name__ == '__main__':
    app.run(port=8060)