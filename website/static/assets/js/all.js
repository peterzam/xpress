(() => {
    "use strict";

    feather.replace({ "aria-hidden": "true" });

    // Graphs
    const ctx = document.getElementById("myChart");
    // eslint-disable-next-line no-unused-vars
    const myChart = new Chart(ctx, {
        type: "line",
        data: {
            labels: [
                "Sunday",
                "Monday",
                "Tuesday",
                "Wednesday",
                "Thursday",
                "Friday",
                "Saturday",
            ],
            datasets: [
                {
                    data: [15, 42, 36, 48, 46, 48, 24],
                    lineTension: 0,
                    backgroundColor: "transparent",
                    borderColor: "#007bff",
                    borderWidth: 4,
                    pointBackgroundColor: "#007bff",
                    color: "#FFFFFF",
                },
            ],
        },
        options: {
            scales: {
                yAxes: [
                    {
                        ticks: {
                            fontColor: "white",
                            beginAtZero: false,
                        },
                        gridLines: {
                            color:'white'
                        },
                    },
                ],
                xAxes: [
                    {
                        ticks: {
                            fontColor: "white",
                            beginAtZero: false,
                        },
                        gridLines: {
                            display: false
                        },
                    },
                ],
            },

            legend: {
                display: false,
            },
        },
    });
})();
