const from = document.querySelector('#from');
const to = document.querySelector('#to');
const iterations = document.querySelector('#iterations');
const dots = document.querySelector('#dots');

const button = document.querySelector('#form-button');
const canvasCtx = document.querySelector('#chart').getContext('2d');

let myChart;

try {
    const rangeRegex = /^(?:-?\d+(?:\.\d+)?|-?(?:\d+(?:\.\d+)?)?pi)$/i;

    button.addEventListener('click', async (event) => {
        if (!rangeRegex.test(from.value) || !rangeRegex.test(to.value)) throw new Error('Invalid input!')

        const response = await fetch(`http://localhost:3007/lab1?from=${from.value}&to=${to.value}&iterations=${iterations.value}&dots=${dots.value}`);

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
            myChart.update();
            return true;
        }

        myChart = new Chart(canvasCtx, {
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
                    }
                ]
            }
        });

        return true;
    }
} catch (error) {
    alert(error.message);
    console.error(error.message);
}