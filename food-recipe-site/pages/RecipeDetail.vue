<template>
    <section class="recipe-detail py-12">
      <div class="container mx-auto">
        <div class="flex flex-col lg:flex-row">
          <div class="lg:w-1/2">
            <img src="/images/recipe-detail.jpg" alt="Recipe" class="w-full rounded-lg">
          </div>
          <div class="lg:w-1/2 lg:pl-12 mt-8 lg:mt-0">
            <h1 class="text-3xl font-bold mb-4">Recipe Title</h1>
            <p class="text-gray-700 mb-6">Detailed description about the recipe...</p>
  
            <div class="mb-8">
              <h2 class="text-xl font-semibold mb-4">Ingredients:</h2>
              <ul class="list-disc pl-6">
                <li>Ingredient 1</li>
                <li>Ingredient 2</li>
                <li>Ingredient 3</li>
              </ul>
            </div>
  
            <div class="mb-8">
              <h2 class="text-xl font-semibold mb-4">Preparation Steps:</h2>
              <ol class="list-decimal pl-6">
                <li>Step 1: Do this...</li>
                <li>Step 2: Do that...</li>
                <li>Step 3: Enjoy your dish!</li>
              </ol>
            </div>
          </div>
        </div>
      </div>
    </section>
  </template>
  