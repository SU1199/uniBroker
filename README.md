# uniBroker

![Untitled_design-removebg-preview](https://user-images.githubusercontent.com/20323373/168052359-fdb3cbd0-fa43-4f36-85c0-bdb714207f62.png)

This project aims to unify the API of all major Indian brokers to create unified trading API interface.

Most of the professional traders and automated trading systems use multiple brokers to execute and manage their trades, positions and portfolios. Which creates a mess of API's web-apps , logins, 2FA keys and more.

With uniBroker my goal is to create a high performance wrapper in GoLang which can execute and manage trades on multiple platforms **without using their paid proprietary API's** and Instead relying on the API endpoints used by broker's web/mobile app.

**Current version (v0.1) supports the following platforms ->**

- [Zerodha](https://zerodha.com/)
- [Upstox](https://upstox.com/)

**With the support for the following currently up in the pipeline ->**

- [Angel Broking](https://www.angelone.in/)
- [5 Paisa](https://www.5paisa.com/)
- [Finvasia](https://www.finvasia.com/)
- [Fyers](https://fyers.in/)
- [HDFC Securities](https://www.hdfcsec.com/)
- [SBI Securities](https://www.sbisecurities.in/)

I've used burpsuite to reverse engineer the login and flow, and api endpoints for orders, cancellation, margins, order-book and positions.

### Things currently in the works ->

- [ ] More broker support.
- [ ] A web-server which can concurrently execute trades with multiple brokers.
- [ ] A UI layer (see mockups) to make interacting with the server easy.
- [ ] Mutual Funds and ETF support.
- [ ] A JSON parser to parse order data from a JSON file.
- [ ] Support fro live data streaming from multiple brokers and news aggregators.

I've decided to keep the project's source open with an MIT licence as making a this a commercial project is beyond SEBI's regulation and anyone's risk tolerance. Anyone who's interested in contributing to the project can submit a PR or email me at hello@danishjoshi.com .
