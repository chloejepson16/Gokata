// Function to generate a random color
function getRandomColor() {
  const r = Math.floor(Math.random() * 255);
  const g = Math.floor(Math.random() * 255);
  const b = Math.floor(Math.random() * 255);
  return `rgba(${r}, ${g}, ${b}, 0.6)`; // Adjust the alpha value (transparency) as needed
}

// Function to generate an array of random colors for each item
function generateColors(count) {
  const colors = [];
  for (let i = 0; i < count; i++) {
    colors.push(getRandomColor());
  }
  return colors;
}

document
  .getElementById("groceryForm")
  .addEventListener("submit", async (event) => {
    event.preventDefault();

    const id = document.getElementById("id").value;
    const name = document.getElementById("name").value;
    const category = document.getElementById("category").value;
    const price = document.getElementById("price").value;

    const data = {
      ID: id,
      name: name,
      category: category,
      price: price,
    };

    try {
      const response = await fetch(
        "http://localhost:3000/groceries/v2/groceryToDB",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(data),
        }
      );

      if (response.ok) {
        alert("Item added successfully");
        document.getElementById("groceryForm").reset();
      } else {
        alert("failed to add item");
      }
    } catch (error) {
      console.error("Error:", error);
      alert("An error occurred. Please check the console for details.");
    }
  });

async function populateChart() {
  try {
    const response = await fetch(
      "http://localhost:3000/groceries/v2/groceryFromDB"
    );
    const data = await response.json();
    console.log(data);

    const labels = data.map((item) => item.name);
    const prices = data.map((item) => parseFloat(item.price.replace("$", "")));
    const totalPrice = prices.reduce((sum, price) => sum + price, 0);

    const backgroundColors = generateColors(labels.length);
    const borderColors = backgroundColors.map((color) =>
      color.replace("0.6", "1")
    );

    const ctx = document.getElementById("groceryChart").getContext("2d");
    new Chart(ctx, {
      type: "bar",
      data: {
        labels: labels,
        datasets: [
          {
            label: "Price $",
            data: prices,
            backgroundColor: "rgba(75, 192, 192, 0.2)",
            borderColor: "rgba(75, 192, 192, 1)",
            borderWidth: 1,
          },
        ],
      },
      options: {
        scales: {
          y: {
            beginAtZero: true,
          },
        },
      },
    });
    const pieCtx = document
      .getElementById("priceBreakdownChart")
      .getContext("2d");
    new Chart(pieCtx, {
      type: "pie",
      data: {
        labels: labels,
        datasets: [
          {
            label: "Price Breakdown",
            data: prices,
            backgroundColor: backgroundColors,
            borderColor: borderColors,
            borderWidth: 1,
          },
        ],
      },
      options: {
        plugins: {
          tooltip: {
            callbacks: {
              label: function (tooltipItem) {
                const price = prices[tooltipItem.dataIndex];
                const percentage = ((price / totalPrice) * 100).toFixed(2);
                return `$${price.toFixed(2)} (${percentage}%)`;
              },
            },
          },
          datalabels: {
            color: "#fff", // Customize the text color
            formatter: (value, context) => {
              const index = context.dataIndex;
              const price = prices[index];
              return `$${price.toFixed(2)}`;
            },
            anchor: "end",
            align: "end",
          },
        },
      },
    });
  } catch (error) {
    console.error("Error fetching or processing data:", error);
  }
}

populateChart();
