package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"terraform-provider-cocktails/model"
)

func resourceCocktail() *schema.Resource {
	return &schema.Resource{
		Create: resourceCocktailCreate,
		Read:   resourceCocktailRead,
		Update: resourceCocktailUpdate,
		Delete: resourceCocktailDelete,
		// CRUD-Funktionen

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"instruction": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"amount_cl": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"ingredient_id": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceCocktailCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	// Hier rufst du deine REST-API auf, um einen Cocktail anzulegen.
	// Für dieses Beispiel setzen wir einfach den Namen als ID.
	d.SetId(name)
	return nil
}

func resourceCocktailRead(d *schema.ResourceData, m interface{}) error {
	// Hier würdest du den aktuellen Cocktail-Status über die API abfragen.
	return nil
}

func resourceCocktailUpdate(d *schema.ResourceData, m interface{}) error {
	// API-Aufruf zum Aktualisieren der Cocktail-Daten
	return nil
}

func resourceCocktailDelete(d *schema.ResourceData, m interface{}) error {
	// API-Aufruf zum Löschen des Cocktails
	d.SetId("")
	return nil
}

func fromCocktail(d *schema.ResourceData, cocktail *model.Cocktail) error {
	d.SetId(strconv.Itoa(cocktail.CocktailID))

	if err := d.Set("name", cocktail.Name); err != nil {
		return err
	}

	instructionsRaw := make([]any, len(cocktail.Instructions))

	for i, instruction := range cocktail.Instructions {
		instructionsRaw[i] = map[string]any{
			"amount_cl":     instruction.AmountCL,
			"ingredient_id": fromID(instruction),
		}
	}

	if err := d.Set("instruction", instructionsRaw); err != nil {
		return err
	}

	return nil
}

func toCocktail(d *schema.ResourceData) *model.Cocktail {
	name := d.Get("name").(string)
	instructionsRaw := d.Get("instruction").([]any)

	instructions := make([]*model.Instruction, len(instructionsRaw))
	for i, instructionRaw := range instructionsRaw {
		instruction := instructionRaw.(map[string]any)
		instructions[i] = &model.Instruction{
			AmountCL:     instruction["amount_cl"].(int),
			IngredientID: toID(instruction["ingredient_id"].(string)),
		}
	}

	return &model.Cocktail{
		Name:         name,
		Instructions: instructions,
	}
}

func fromID(instruction *model.Instruction) string {
	return strconv.Itoa(instruction.IngredientID)
}

func toID(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}
