package services

import (
	"context"
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "strings"

    "github.com/google/generative-ai-go/genai"
    "google.golang.org/api/option"
)

// IngredientMap represents the mapping of meals to their ingredients.
type IngredientMap map[string][]string

// AIAgent is the struct that holds the logic for processing meals and generating grocery lists.
type AIAgent struct{}

// NewAIAgent creates a new instance of AIAgent.
func NewAIAgent() *AIAgent {
	return &AIAgent{}
}

// ExpandMeals takes a list of meals and returns a map of ingredients needed for each meal.
func (a *AIAgent) ExpandMeals(meals []string) (IngredientMap, error) {
	if len(meals) == 0 {
		return nil, errors.New("no meals provided")
	}

	// Simulated ingredient expansion logic
	ingredients := make(IngredientMap)
	for _, meal := range meals {
		switch strings.ToLower(meal) {
		case "pasta":
			ingredients["pasta"] = []string{"spaghetti", "olive oil", "garlic", "tomatoes"}
		case "salad":
			ingredients["salad"] = []string{"lettuce", "tomatoes", "cucumber", "olive oil"}
		case "pancakes":
			ingredients["pancakes"] = []string{"flour", "milk", "eggs", "baking powder"}
		default:
			return nil, errors.New("unknown meal: " + meal)
		}
	}
	return ingredients, nil
}

// MergeIngredients combines the ingredient lists from multiple meals, removing duplicates.
func (a *AIAgent) MergeIngredients(ingredientMap IngredientMap) []string {
	ingredientSet := make(map[string]struct{})
	for _, ingredients := range ingredientMap {
		for _, ingredient := range ingredients {
			ingredientSet[ingredient] = struct{}{}
		}
	}

	var mergedIngredients []string
	for ingredient := range ingredientSet {
		mergedIngredients = append(mergedIngredients, ingredient)
	}
	return mergedIngredients
}

// SuggestHealthierSwaps suggests healthier alternatives for the given ingredient list.
func (a *AIAgent) SuggestHealthierSwaps(ingredients []string) (map[string]string, error) {
    swaps := make(map[string]string)
    for _, ingredient := range ingredients {
        normalizedIngredient := strings.ToLower(strings.TrimSpace(ingredient))
        switch normalizedIngredient {
        // Flours and grains
        case "flour", "white flour", "all-purpose flour":
            swaps[ingredient] = "almond flour or whole wheat flour"
        case "white rice":
            swaps[ingredient] = "brown rice or quinoa"
        case "pasta", "white pasta":
            swaps[ingredient] = "whole grain pasta or zucchini noodles"
        case "bread", "white bread":
            swaps[ingredient] = "whole grain bread or ezekiel bread"
        case "breadcrumbs":
            swaps[ingredient] = "almond meal or ground flaxseed"
            
        // Dairy products	
        case "milk", "whole milk":
            swaps[ingredient] = "oat milk or almond milk"
        case "cream", "heavy cream":
            swaps[ingredient] = "coconut cream or cashew cream"
        case "sour cream":
            swaps[ingredient] = "greek yogurt"
        case "butter":
            swaps[ingredient] = "ghee or coconut oil"
        case "cheese", "cheddar cheese":
            swaps[ingredient] = "nutritional yeast or cashew cheese"
        case "ice cream":
            swaps[ingredient] = "banana nice cream or coconut milk ice cream"
            
        // Oils and fats	
        case "vegetable oil", "canola oil":
            swaps[ingredient] = "avocado oil or olive oil"
        case "olive oil":
            swaps[ingredient] = "avocado oil"
        case "margarine":
            swaps[ingredient] = "coconut oil or avocado oil spread"
            
        // Sweeteners	
        case "sugar", "white sugar":
            swaps[ingredient] = "maple syrup, honey, or monk fruit sweetener"
        case "brown sugar":
            swaps[ingredient] = "coconut sugar"
        case "corn syrup":
            swaps[ingredient] = "maple syrup or date syrup"
            
        // Proteins	
        case "ground beef":
            swaps[ingredient] = "lean ground turkey or plant-based ground"
        case "bacon":
            swaps[ingredient] = "turkey bacon or tempeh bacon"
        case "chicken", "chicken breast":
            swaps[ingredient] = "organic free-range chicken"
        case "tuna", "canned tuna":
            swaps[ingredient] = "wild-caught salmon"
            
        // Snacks and processed foods	
        case "potato chips":
            swaps[ingredient] = "baked vegetable chips or air-popped popcorn"
        case "crackers", "white crackers":
            swaps[ingredient] = "whole grain crackers or seed crackers"
        case "chocolate", "milk chocolate":
            swaps[ingredient] = "dark chocolate (70%+ cacao)"
        case "candy":
            swaps[ingredient] = "dried fruit or dark chocolate"
            
        // Condiments and sauces	
        case "mayonnaise", "mayo":
            swaps[ingredient] = "avocado-based mayo or greek yogurt"
        case "ketchup":
            swaps[ingredient] = "tomato paste or sugar-free ketchup"
        case "ranch dressing":
            swaps[ingredient] = "greek yogurt-based dressing"
        case "soy sauce":
            swaps[ingredient] = "coconut aminos or tamari"
        case "salt", "table salt":
            swaps[ingredient] = "himalayan pink salt or sea salt"
            
        // Baking ingredients	
        case "baking powder":
            swaps[ingredient] = "aluminum-free baking powder"
        case "chocolate chips":
            swaps[ingredient] = "cacao nibs or dark chocolate chips"
            
        // Beverages	
        case "soda", "cola":
            swaps[ingredient] = "sparkling water with fresh fruit"
        case "fruit juice":
            swaps[ingredient] = "whole fruits or vegetable juice"
        case "coffee creamer":
            swaps[ingredient] = "almond milk or coconut milk"
        }
    }
    return swaps, nil
}

// ToJSON converts the ingredient map to a JSON string.
func (a *AIAgent) ToJSON(ingredientMap IngredientMap) (string, error) {
	data, err := json.Marshal(ingredientMap)
	if (err != nil) {
		return "", err
	}
	return string(data), nil
}

func GenerateGroceryList(meals []string) (map[string]interface{}, error) {
    if len(meals) == 0 {
        return nil, errors.New("no meals provided")
    }

    // Get the Gemini API key from the environment
    apiKey := os.Getenv("GEMINI_API_KEY")
    if apiKey == "" {
        return nil, errors.New("Gemini API key not found in environment variables")
    }

    // Initialize the Gemini client
    ctx := context.Background()
    client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
    if err != nil {
        return nil, fmt.Errorf("failed to initialize Gemini client: %v", err)
    }
    defer client.Close()

    // Select the Gemini model
    model := client.GenerativeModel("gemini-1.5-pro")

    // Updated prompt that explicitly requests JSON format
    prompt := fmt.Sprintf(`Given the following meals: %s
Generate a detailed grocery shopping list.
Group items by categories like "Vegetables", "Dairy", "Proteins", etc.
Respond ONLY with JSON like:

{
  "Vegetables": ["Tomatoes", "Onions"],
  "Dairy": ["Cheese", "Milk"],
  "Proteins": ["Chicken Breast", "Eggs"],
  ...
}

Do NOT include any extra commentary, bullet points, or markdown formatting. Just pure JSON.`, strings.Join(meals, ", "))

    fmt.Printf("Sending prompt to Gemini: %s\n", prompt)

    // Generate content
    resp, err := model.GenerateContent(ctx, genai.Text(prompt))
    if err != nil {
        return nil, fmt.Errorf("Gemini content generation failed: %v", err)
    }

    if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
        return nil, errors.New("no response generated from Gemini")
    }

    // Extract the text response
    responseText, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
    if !ok {
        return nil, errors.New("unexpected response format from Gemini")
    }

    fmt.Printf("Raw response from Gemini: %s\n", string(responseText))

    // Try to clean the response text if it's not properly formatted
    cleanedText := strings.TrimSpace(string(responseText))
    
    // Find the first '{' and last '}' to extract just the JSON part
    startIdx := strings.Index(cleanedText, "{")
    endIdx := strings.LastIndex(cleanedText, "}")
    
    if startIdx >= 0 && endIdx > startIdx {
        cleanedText = cleanedText[startIdx:endIdx+1]
    }

    // Parse the response as JSON
    var groceryList map[string][]string
    if err := json.Unmarshal([]byte(cleanedText), &groceryList); err != nil {
        return nil, fmt.Errorf("failed to parse JSON response: %v (received: %s)", err, cleanedText)
    }

    // Extract all ingredients into a flat list
    var allIngredients []string
    for _, ingredients := range groceryList {
        allIngredients = append(allIngredients, ingredients...)
    }

    // Create an AIAgent to suggest healthier swaps
    agent := NewAIAgent()
    healthierSwaps, err := agent.SuggestHealthierSwaps(allIngredients)
    if err != nil {
        return nil, fmt.Errorf("failed to suggest healthier swaps: %v", err)
    }

    fmt.Printf("Ingredients for healthier swaps: %v\n", allIngredients)
    fmt.Printf("Healthier swaps found: %v\n", healthierSwaps)

    // Create the final response with both grocery list and healthier swaps
    result := map[string]interface{}{
        "grocery_list": groceryList,
        "healthier_swaps": healthierSwaps,
    }

    return result, nil
}


func formatMeals(meals []string) string {
	return "\"" + strings.Join(meals, "\", \"") + "\""
}