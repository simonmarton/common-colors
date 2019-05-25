// eslint-disable-next-line no-unused-vars
const appendUploadButton = (afterUpload = () => {}, getConfig = () => {}, withSteps) => {
  let input = document.querySelector('input[type=file]');
  if (!input) {
    input = document.createElement('input');
    input.setAttribute('type', 'file');
    document.body.prepend(input);
  }

  input.addEventListener('change', evt => {
    evt.preventDefault();

    const {
      files: [image]
    } = input;

    uploadImage(image, getConfig(), withSteps).then(afterUpload);

    return false;
  });
};

// eslint-disable-next-line no-unused-vars
const uploadImage = async (image, config, withSteps) => {
  if (!image) {
    throw new Error('Select an image');
  }

  if (!config) {
    console.log('Using default config');

    config = {
      algorithm: 'yiq',
      transparencyTreshold: 10,
      iterationCount: 3,
      minLuminance: 0.3,
      maxLuminance: 0.9,
      distanceThreshold: 20,
      minSaturation: 0.3
    };
  }

  const formData = new FormData();
  formData.append('image', image);
  formData.append('config', JSON.stringify(config));

  const reader = new FileReader();

  reader.onload = e => {
    const preview = document.querySelector('#preview');
    preview && preview.setAttribute('src', e.target.result);

    const icon = document.querySelector('#icon');
    icon && icon.setAttribute('src', e.target.result);
  };

  reader.readAsDataURL(image);

  const result = await fetch(`/api/upload${withSteps ? '?steps' : ''}`, {
    method: 'post',
    body: formData
  }).then(res => res.json());

  const { colors, gradient, steps } = result || {};

  return { colors, gradient, steps };
};
