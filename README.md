# Smart Grocery Agent

The **Smart Grocery Agent** is a Go-based application that generates detailed grocery lists for meals provided by the user. It leverages the **Gemini AI API** for dynamic ingredient generation and incorporates an AI agent to enhance the response with healthier ingredient swaps.

## Features

- **Dynamic Grocery List Generation**: Uses the Gemini AI API to generate grocery lists for any meal.
- **Healthier Ingredient Swaps**: Suggests healthier alternatives for a wide range of common ingredients.
- **Customizable Logic**: Includes hardcoded fallback logic for specific meals and ingredients.
- **Error Handling**: Handles API errors and ensures consistent response formatting.

## Project Structure

```
smart-grocery-agent/
├── cmd/
│   └── main.go                # Entry point of the application
├── internal/
│   ├── handlers/
│   │   └── grocery.go         # HTTP handler for the /grocery-list endpoint
│   ├── services/
│   │   └── ai_agent.go        # Core AI agent logic for processing meals and generating grocery lists
├── .env                       # Environment variables (e.g., GEMINI_API_KEY)
├── go.mod                     # Go module file
├── go.sum                     # Dependency lock file
└── README.md                  # Project documentation
```

## Prerequisites

- **Go**: Ensure you have Go installed (version 1.18 or later).
- **Gemini API Key**: Obtain an API key for the Gemini AI API and add it to the `.env` file.

## Installation

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd smart-grocery-agent
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up the `.env` file:
   ```plaintext
   GEMINI_API_KEY=your-gemini-api-key
   ```

## Usage

1. Start the server:

   ```bash
   go run cmd/main.go
   ```

2. Test the `/grocery-list` endpoint using Postman or `curl`:

   - **URL**: `http://localhost:3000/grocery-list`
   - **Method**: `POST`
   - **Headers**:
     - `Content-Type: application/json`
   - **Body**:
     ```json
     {
       "meals": ["Chicken biryani", "sushi", "tacos"]
     }
     ```

3. Example Response:
   ```json
   {
     "grocery_list": {
       "Condiments/Spices": [
         "Biryani Masala",
         "Turmeric",
         "Cumin",
         "Coriander",
         "Chili Powder",
         "Salt",
         "Pepper",
         "Soy Sauce",
         "Wasabi",
         "Ginger (Pickled)",
         "Taco Seasoning",
         "Salsa"
       ],
       "Dairy": ["Yogurt (for biryani)", "Sour Cream (for tacos)"],
       "Grains/Starches": ["Basmati Rice", "Sushi Rice", "Taco Shells"],
       "Other": ["Nori Seaweed Sheets", "Lime Wedges"],
       "Proteins": [
         "Chicken Thighs",
         "Sushi-grade Salmon or Tuna",
         "Ground Beef or Chicken"
       ],
       "Vegetables": [
         "Onions",
         "Tomatoes",
         "Cilantro",
         "Lettuce",
         "Avocado",
         "Cucumber",
         "Ginger",
         "Garlic",
         "Green Onions",
         "Jalapenos",
         "Serrano Peppers"
       ]
     },
     "healthier_swaps": {
       "Salt": "himalayan pink salt or sea salt",
       "Soy Sauce": "coconut aminos or tamari"
     }
   }
   ```

## Key Components

### 1. **Gemini AI Integration**

- The application uses the Gemini AI API to dynamically generate grocery lists based on the provided meals.
- The API is accessed via the `genai` Go client library.

### 2. **AI Agent**

- The AI agent (`ai_agent.go`) enhances the Gemini response by:
  - Suggesting healthier ingredient swaps for a wide range of ingredients.
  - Providing fallback logic for predefined meals (e.g., pasta, salad, pancakes).

### 3. **Healthier Swaps**

- The `SuggestHealthierSwaps` function in `ai_agent.go` provides healthier alternatives for ingredients such as:
  - **Flours and Grains**: White flour → Almond flour, White rice → Brown rice
  - **Dairy Products**: Milk → Oat milk, Butter → Coconut oil
  - **Oils and Fats**: Vegetable oil → Avocado oil
  - **Sweeteners**: Sugar → Maple syrup or monk fruit sweetener
  - **Proteins**: Ground beef → Lean ground turkey
  - **Snacks**: Potato chips → Baked vegetable chips
  - **Condiments**: Soy sauce → Coconut aminos

### 4. **Utilities**

- The `MergeIngredients` function in `ai_agent.go` merges multiple ingredient lists and removes duplicates.

## Environment Variables

- **GEMINI_API_KEY**: Your API key for accessing the Gemini AI API. Add this to the `.env` file.

## Dependencies

- [Fiber](https://gofiber.io/): Web framework for building HTTP APIs.
- [Gemini AI Go Client](https://pkg.go.dev/github.com/google/generative-ai-go): Client library for interacting with the Gemini AI API.
- [godotenv](https://github.com/joho/godotenv): For loading environment variables from a `.env` file.

## Future Enhancements

- Add support for dietary restrictions (e.g., vegetarian, gluten-free).
- Integrate pricing and availability data for grocery items.
- Provide recipe suggestions based on the grocery list.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Acknowledgments

- [Google Gemini AI](https://ai.google/): For providing the AI model used in this project.
- [Fiber Framework](https://gofiber.io/): For the HTTP server framework.
- [godotenv](https://github.com/joho/godotenv): For environment variable management.
