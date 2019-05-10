const form = document.querySelector('form#upload');
const input = document.querySelector('input[type=file]');

// form.addEventListener('submit', evt => {
input.addEventListener('change', evt => {
  evt.preventDefault();

  // const { files } = form.querySelector('input[type=file]');
  const { files } = input;

  if (!files.length) {
    return alert('Select an image');
  }

  const img = files[0];
  const formData = new FormData();
  formData.append('image', img);

  const reader = new FileReader();

  reader.onload = e => {
    document.querySelector('#preview').setAttribute('src', e.target.result);
  };

  reader.readAsDataURL(img);

  fetch('/api/upload', {
    method: 'post',
    body: formData
  })
    .then(res => res.json())
    .then(({ colors } = {}) => {
      const result = document.querySelector('#result');
      result.innerHTML = '';

      colors = colors || [];

      const total = colors.reduce((total, { weight }) => total + weight, 0);

      colors.forEach(({ value, weight }) => {
        const percetage = (weight / total) * 100;
        console.log({ value, weight, percetage });

        const container = document.createElement('div');
        if (percetage < 5) {
          container.className = 'light';
        }

        const color = document.createElement('span');
        color.className = 'color';
        color.style.background = value;

        const weightElem = document.createElement('span');
        weightElem.className = 'weight';
        weightElem.innerText = `${percetage.toFixed(2)}% (${weight})`;

        container.appendChild(color);
        container.appendChild(weightElem);

        result.appendChild(container);
      });

      console.log({ colors });
    });

  return false;
});
