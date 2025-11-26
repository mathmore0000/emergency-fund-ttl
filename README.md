# Emergency Fund Time-to-Live Calculator

A small Go program that estimates how many months an emergency fund will last
given an initial balance, a fixed monthly withdrawal, and a monthly return
rate (expressed as a percent, for example pass `1` for 1%).

Why this matters: many people keep their emergency fund in an accessible,
low-risk place (for example: a high-yield savings account, money market
fund, or short-term bond fund) rather than as idle cash. Those vehicles can
generate a small but meaningful return while still allowing quick access to
the money. Because of that common practice, including an expected monthly
return when estimating time-to-live gives a more realistic result.

The calculation models each month as:
    1. Apply monthly return to the current balance
    2. Subtract the fixed monthly withdrawal

The program returns the number of months (including a fractional last month)
until the balance reaches zero or an error if the inputs indicate the fund will
never deplete.

## Usage

Run the program and enter the requested values when prompted:

```sh
go run main.go
```

You will be prompted for:
- Amount saved (initial balance)
- Fixed monthly withdrawal
- Monthly return (as a percent, e.g. `1` for 1%)

Example:

```
Enter the amount saved:
150000
Enter the fixed monthly withdrawal:
10000
Enter the monthly return (as a decimal, e.g. 1 for 1%):
1
It will take 15.34 months to deplete the savings.
```

## Notes and caveats

- The function returns errors for invalid or contradictory inputs (negative
    monthly return, withdrawal that causes non-depletion, etc.).
- The implementation preserves the original calculation logic and reports a
    fractional final month.