// Set up config
document.querySelectorAll('input[type="range"]').forEach(elem => {
  elem.nextElementSibling.value = elem.value;
  elem.oninput = e => {
    e.target.nextElementSibling.value = e.target.value;
  };
});

const getConfig = () =>
  Array.from(document.querySelectorAll('input[type="range"]')).reduce(
    (config, elem) => ({
      ...config,
      [elem.name]: Number(elem.value)
    }),
    { algorithm: document.querySelector('select[name="algorithm"]').value }
  );

const setGradientColors = (c1, c2) => {
  console.log('setting gradient colors', c1, c2);
  document.getElementById('main-color-1').setAttribute('stop-color', c1);
  document.getElementById('main-color-2').setAttribute('stop-color', c2);
};

const handleResult = ({ colors, gradient }) => {
  const result = document.querySelector('#result');
  result.innerHTML = '';

  colors = colors || [];
  gradient = gradient || [];

  setGradientColors(...gradient);

  const total = colors.reduce((total, { weight }) => total + weight, 0);

  colors.forEach(({ value, weight, hueDistance }) => {
    const percetage = (weight / total) * 100;

    const container = document.createElement('div');
    if (percetage < 5) {
      container.className = 'light';
    }

    const color = document.createElement('span');
    color.className = 'color';
    color.style.background = value;

    const weightElem = document.createElement('span');
    weightElem.className = 'weight';
    weightElem.innerText = `${percetage.toFixed(2)}% (h: ${hueDistance.toFixed(2)}, w: ${weight})`;

    container.appendChild(color);
    container.appendChild(weightElem);

    result.appendChild(container);
  });
};

document.getElementById('upload').addEventListener('click', () => {
  let input = document.querySelector('input[type=file]');

  const {
    files: [image]
  } = input;

  uploadImage(image, getConfig()).then(handleResult);
});

appendUploadButton(handleResult, getConfig);
