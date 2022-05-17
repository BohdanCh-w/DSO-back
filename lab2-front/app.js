// Input data
const from = document.querySelector('#from');
const to = document.querySelector('#to');
const iterations = document.querySelector('#iterations');
const dots = document.querySelector('#dots');
const pointsContainer = document.querySelector('.form-container-points');

// Buttons
const button = document.querySelector('#form-button');
const addPointButton = document.getElementById('add-point-button');
const removePointButton = document.getElementById('remove-point-button');

// Canvas
const canvasCtx = document.querySelector('#chart').getContext('2d');
let myChart;

let values = ['2.81', '2.95', '3.21', '3.39', '4.54', '3.76', '4.03', '3.54'];
let currentPointsCount = 6;

try {
    drawDots();

    const rangeRegex = /^(?:-?\d+(?:\.\d+)?|-?(?:\d+(?:\.\d+)?)?pi)$/i;

    button.addEventListener('click', async (event) => {
        if (!rangeRegex.test(from.value) || !rangeRegex.test(to.value)) throw new Error('Invalid input!');

        values = [];
        document.querySelectorAll('.form-item-point').forEach(item => values.push(item.value));

        console.log(values);

        const response = await fetch(`http://localhost:3007/lab2?from=${from.value}&to=${to.value}&dots=${dots.value}&points=${values.join(',')}`);

        console.log(response);
        const result = await response.json();
        console.log(result);


        if (!response.ok) {
            console.error(`Error: status ${response.status} ${response.statusText}, ${result.error}`)
            return
        }
        drawFunctions(result);
    });

    function drawFunctions(array = []) {
        if (myChart) {
            myChart.data.labels = array.map(element => element.x.toFixed(3));
            myChart.data.datasets[0].data = array.map(element => element.y.toFixed(3));
            myChart.data.datasets[1].data = array.map(element => element.yf.toFixed(3));
            myChart.data.datasets[2].data = array.map(element => element.ys.toFixed(3));
            myChart.update();
            return true;
        }

        myChart = new Chart(canvasCtx, {
            options: {
                responsive: true,
                maintainAspectRatio: false,
                scales: {
                    yAxes: [{
                        ticks: {
                            beginAtZero: true
                        }
                    }]
                }
            },
            type: 'line',
            data: {
                labels: array.map(element => element.x.toFixed(3)),
                datasets: [
                    {
                        label: 'Function',
                        data: array.map(element => element.y),
                        backgroundColor: 'rgba(0, 0, 255, 0.2)',
                        borderColor: 'rgba(0, 0, 0, 1)',
                        pointStyle: 'dash',
                    },
                    {
                        label: 'Fourier',
                        data: array.map(element => element.yf),
                        backgroundColor: 'rgba(255, 255, 0, 0.2)',
                        borderColor: 'rgba(0, 180, 180, 1)',
                        pointStyle: 'dash',
                    },
                    {
                        label: 'Quadratic',
                        data: array.map(element => element.ys),
                        backgroundColor: 'rgba(255, 0, 255, 0.2)',
                        borderColor: 'rgba(180, 180, 0, 1)',
                        pointStyle: 'dash',
                    }
                ]
            }
        });

        return true;
    }

    function drawDots(count = 6) {
        pointsContainer.innerHTML = '';

        for (let i = 0; i < count; i++) {
            const dot = document.createElement('input');
            dot.classList.add('form-item', 'form-item-point');
            dot.type = 'number';

            dot.value = i < 6 ? values[i] : 0;

            pointsContainer.appendChild(dot);
        }
    }
} catch (error) {
    alert(error.message);
    console.error(error.message);
}

// document.querySelectorAll('.form-item-point').forEach(item => item.addEventListener('click', event => event.innerHTML = ''));

addPointButton.addEventListener('click', () => drawDots(++currentPointsCount));

removePointButton.addEventListener('click', () => currentPointsCount > 3 ? drawDots(--currentPointsCount) : alert("Number of points shouldn't be less than 3"));