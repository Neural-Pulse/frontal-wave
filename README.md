# Next Best Offer (NBO) in Go

## Introduction

Next Best Offer (NBO) is a recommendation technique used by many companies to suggest products or services to customers based on their preferences and purchase history. This project implements a simple NBO system in Go, which reads user data and preferences from a CSV file and generates next best offer recommendations for each user.

## Theory of Next Best Offer (NBO)

The idea behind Next Best Offer is to offer the customer the best possible offer considering their preferences and past behaviors. This can increase customer satisfaction, improve conversion rates, and boost sales.

The general process of NBO involves:

1. **Data Collection**: Collecting customer data, including their purchase history, past interactions, and stated preferences.
2. **Data Analysis**: Analyzing the collected data to understand customer preferences and behaviors.
3. **Recommendation Generation**: Based on the data analysis, generating personalized recommendations for each customer.
4. **Presentation of Recommendations**: Presenting the recommendations to customers through various channels such as websites, mobile apps, emails, etc.

## Input Data

The NBO system needs to receive user data and their preferences in a structured format. The input data should be provided in a CSV file with the following columns:

- **id_usuario**: A unique identifier for each user.
- **produto**: The name or identifier of the product.
- **pontuacao_preferencia**: A score indicating the user's preference for the product. The higher the score, the higher the preference.

Here's an example of how the data should be structured:

```
id_usuario,produto,pontuacao_preferencia
1,ProductA,5
1,ProductB,3
1,ProductC,2
2,ProductA,4
2,ProductB,4
2,ProductC,3
3,ProductA,2
3,ProductB,3
3,ProductC,5
4,ProductA,3
4,ProductB,2
4,ProductC,4
```

## How to Run the Project

1. **Install Go**: If you haven't already installed Go, follow the instructions at [https://golang.org/doc/install](https://golang.org/doc/install) to install Go on your system.
2. **Clone the Repository**: Clone this repository to your local machine.
3. **Run the Program**: Run the main program `main.go` providing the path to the data CSV file as an argument. For example:

    ```sh
    go run main.go data.csv
    ```

    Make sure to replace `"data.csv"` with the actual path to your data CSV file.

## Contributions

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.