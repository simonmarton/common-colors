/* global THREE */

const v3 = (...args) => new THREE.Vector3(...args);

class Point {
  constructor(position, radius, color) {
    // this.birth = Date.now();
    this.growTimeout = Date.now() + Math.random() * 500;
    this.scale = 0.01;
    this.geometry = new THREE.SphereGeometry(radius, 32, 32);

    // this.material = new THREE.MeshBasicMaterial({ color });
    this.material = new THREE.MeshPhongMaterial({ color, shininess: 5 });
    this.mesh = new THREE.Mesh(this.geometry, this.material);

    this.mesh.scale.set(v3(this.scale));

    this.mesh.position.set(position.x, position.y, position.z);
  }

  static FromColor({ r, g, b, weight }) {
    const w = Math.min(Math.max(5, weight / 20), 50);
    return new Point(new THREE.Vector3(r, g, b), w, `rgb(${r}, ${g}, ${b})`);
  }

  update() {
    if (this.scale < 1 && this.growTimeout < Date.now()) {
      this.scale += 0.05;
      const scale = this.scale * (2 - this.scale);

      this.mesh.scale.x = scale;
      this.mesh.scale.y = scale;
      this.mesh.scale.z = scale;
    }
  }

  getMesh() {
    return this.mesh;
  }
}

class Line {
  constructor(startPoint, endPoint, startColor, endColor) {
    const geometry = new THREE.Geometry();

    geometry.vertices = [startPoint, endPoint];
    geometry.colors = [startColor, endColor];

    const material = new THREE.LineBasicMaterial({
      color: 'white',
      vertexColors: THREE.VertexColors
    });

    this.mesh = new THREE.LineSegments(geometry, material);
  }

  getMesh() {
    return this.mesh;
  }
}

class Application {
  constructor() {
    this.objects = [];

    this.createScene();
  }

  createScene() {
    this.scene = new THREE.Scene();
    this.scene.background = new THREE.Color('#333');
    this.camera = new THREE.PerspectiveCamera(40, window.innerWidth / window.innerHeight, 1, 2000);
    this.camera.position.set(500, 450, 450);

    this.controls = new THREE.OrbitControls(this.camera);
    this.controls.target.set(128, 128, 128);
    this.controls.enableZoom = false;
    this.controls.minDistance = 800;
    this.controls.maxDistance = 800;

    this.controls.update();
    this.controls.autoRotate = true;
    this.controls.autoRotate = true;
    this.controls.autoRotateSpeed = 1.5;

    this.renderer = new THREE.WebGLRenderer();
    this.renderer.setSize(window.innerWidth, window.innerHeight);

    window.onresize = () => {
      this.renderer.setSize(window.innerWidth, window.innerHeight);
      this.camera.aspect = window.innerWidth / window.innerHeight;
      this.camera.updateProjectionMatrix();
    };
    document.body.appendChild(this.renderer.domElement);

    const particleLight = new THREE.Mesh(
      new THREE.SphereBufferGeometry(3, 8, 8),
      new THREE.MeshBasicMaterial({ color: 0xffffff })
    );
    particleLight.position.set(2, 5, 3);
    var directionalLight = new THREE.DirectionalLight('#fff', 0.5);
    directionalLight.position.set(1, 1, 1).normalize();
    this.scene.add(directionalLight);
    this.scene.add(new THREE.AmbientLight('#aaa'));

    this.createAxis();

    this.render();
  }

  createAxis() {
    const add = (r1, g1, b1, r2, g2, b2) =>
      this.scene.add(
        new Line(
          v3(r1, g1, b1),
          v3(r2, g2, b2),
          new THREE.Color(r1 / 255, g1 / 255, b1 / 255),
          new THREE.Color(r2 / 255, g2 / 255, b2 / 255)
        ).getMesh()
      );

    add(0, 0, 0, 255, 0, 0);
    add(0, 0, 0, 0, 255, 0);
    add(0, 0, 0, 0, 0, 255);

    add(255, 255, 0, 255, 0, 0);
    add(255, 255, 0, 0, 255, 0);
    add(255, 255, 0, 255, 255, 255);

    add(0, 255, 255, 0, 255, 0);
    add(0, 255, 255, 0, 0, 255);
    add(0, 255, 255, 255, 255, 255);

    add(255, 0, 255, 0, 0, 255);
    add(255, 0, 255, 255, 0, 0);
    add(255, 0, 255, 255, 255, 255);
  }

  render() {
    requestAnimationFrame(() => {
      this.render();
      this.controls.update();
    });

    this.objects.forEach(object => {
      object.update && object.update();
    });

    this.renderer.render(this.scene, this.camera);
  }

  add(mesh) {
    this.objects.push(mesh);
    this.scene.add(mesh.getMesh());
  }

  clear() {
    this.objects.forEach(obj => {
      const mesh = obj.getMesh();
      this.scene.remove(mesh);
      mesh.geometry.dispose();
      mesh.material.dispose();
    });

    this.objects = [];
  }
}

const app = new Application();

let allSteps;
let stepIdx = 0;

const next = document.getElementById('next');

// eslint-disable-next-line no-undef
appendUploadButton(
  ({ steps }) => {
    allSteps = steps;
    stepIdx = 0;

    next.removeAttribute('disabled');

    process();
  },
  () => ({
    algorithm: 'yiq',
    distanceThreshold: 30
  }),
  true
);

const process = () => {
  if (stepIdx == allSteps.length) {
    return;
  }

  next.innerText = `next ${stepIdx + 1} / ${allSteps.length}`;

  app.clear();

  const step = allSteps[stepIdx++];

  step.forEach(point => {
    if (point.weight < 5) return;
    app.add(Point.FromColor(point));
  });

  if (stepIdx === allSteps.length) {
    next.setAttribute('disabled', true);
  }
};

document.getElementById('next').onclick = () => {
  process();
};

// eslint-disable-next-line no-unused-vars
const mock = () => {
  // prettier-ignore
  allSteps = [
    [{r:245,g:233,b:36,weight:123},{r:250,g:200,b:27,weight:110},{r:28,g:165,b:221,weight:106},{r:164,g:54,b:138,weight:105},{r:27,g:116,b:181,weight:84},{r:252,g:226,b:25,weight:80},{r:65,g:177,b:168,weight:74},{r:137,g:193,b:63,weight:71},{r:248,g:184,b:27,weight:69},{r:180,g:51,b:119,weight:59},{r:26,g:134,b:196,weight:47},{r:166,g:204,b:53,weight:46},{r:77,g:94,b:165,weight:43},{r:244,g:157,b:30,weight:43},{r:197,g:44,b:89,weight:40},{r:132,g:70,b:149,weight:39},{r:249,g:237,b:93,weight:37},{r:46,g:171,b:201,weight:35},{r:122,g:188,b:73,weight:34},{r:240,g:129,b:33,weight:34},{r:85,g:181,b:134,weight:34},{r:106,g:186,b:102,weight:34},{r:54,g:104,b:171,weight:33},{r:34,g:147,b:204,weight:32},{r:236,g:98,b:37,weight:31},{r:195,g:216,b:48,weight:30},{r:104,g:80,b:156,weight:28},{r:249,g:205,b:72,weight:27},{r:170,g:215,b:228,weight:27},{r:65,g:179,b:223,weight:26},{r:231,g:69,b:41,weight:25},{r:66,g:92,b:166,weight:25},{r:226,g:42,b:44,weight:25},{r:213,g:223,b:37,weight:22},{r:214,g:40,b:62,weight:21},{r:42,g:126,b:186,weight:20},{r:250,g:215,b:121,weight:20},{r:168,g:208,b:96,weight:17},{r:234,g:232,b:145,weight:15},{r:139,g:205,b:175,weight:14},{r:102,g:190,b:166,weight:14},{r:193,g:106,b:161,weight:13},{r:194,g:222,b:132,weight:13},{r:122,g:195,b:216,weight:13},{r:146,g:201,b:110,weight:13},{r:199,g:230,b:233,weight:12},{r:96,g:194,b:230,weight:11},{r:172,g:69,b:147,weight:11},{r:205,g:132,b:176,weight:11},{r:242,g:155,b:91,weight:11},{r:180,g:88,b:157,weight:11},{r:216,g:153,b:185,weight:10},{r:93,g:181,b:202,weight:10},{r:247,g:193,b:132,weight:9},{r:88,g:160,b:205,weight:9},{r:18,g:157,b:217,weight:9},{r:194,g:228,b:191,weight:9},{r:246,g:178,b:67,weight:9},{r:70,g:142,b:194,weight:9},{r:239,g:126,b:87,weight:8},{r:120,g:137,b:188,weight:8},{r:142,g:161,b:202,weight:8},{r:138,g:207,b:234,weight:8},{r:211,g:95,b:126,weight:7},{r:214,g:56,b:80,weight:7},{r:101,g:130,b:185,weight:7},{r:226,g:178,b:202,weight:7},{r:146,g:185,b:216,weight:7},{r:191,g:81,b:139,weight:6},{r:100,g:118,b:179,weight:6},{r:78,g:113,b:176,weight:6},{r:214,g:234,b:180,weight:6},{r:143,g:57,b:143,weight:6},{r:159,g:37,b:129,weight:6},{r:244,g:211,b:215,weight:6},{r:206,g:69,b:107,weight:5},{r:172,g:214,b:141,weight:5},{r:224,g:120,b:137,weight:5},{r:245,g:199,b:185,weight:5},{r:171,g:179,b:211,weight:5},{r:171,g:218,b:197,weight:5},{r:115,g:172,b:210,weight:5},{r:119,g:151,b:197,weight:5},{r:234,g:99,b:89,weight:4},{r:192,g:123,b:176,weight:4},{r:211,g:172,b:207,weight:4},{r:155,g:80,b:153,weight:3},{r:244,g:168,b:134,weight:3},{r:177,g:35,b:104,weight:3},{r:230,g:78,b:83,weight:3},{r:237,g:147,b:149,weight:3},{r:195,g:202,b:226,weight:3},{r:9,g:108,b:175,weight:3},{r:34,g:96,b:167,weight:2},{r:126,g:194,b:115,weight:2},{r:206,g:224,b:90,weight:2},{r:241,g:174,b:170,weight:2},{r:228,g:201,b:218,weight:1},{r:228,g:235,b:92,weight:1},{r:204,g:153,b:198,weight:1},{r:118,g:59,b:141,weight:1},{r:93,g:68,b:149,weight:1},{r:116,g:202,b:235,weight:1}],
  [{r:246,g:220,b:43,weight:420},{r:34,g:158,b:211,weight:255},{r:163,g:58,b:136,weight:250},{r:138,g:195,b:74,weight:217},{r:77,g:176,b:167,weight:155},{r:35,g:114,b:179,weight:142},{r:246,g:172,b:36,weight:132},{r:82,g:92,b:164,weight:109},{r:209,g:44,b:72,weight:98},{r:167,g:211,b:213,weight:90},{r:237,g:114,b:43,weight:77},{r:222,g:222,b:159,weight:45},{r:203,g:114,b:157,weight:40},{r:118,g:186,b:217,weight:33},{r:195,g:216,b:50,weight:32},{r:223,g:163,b:182,weight:30},{r:230,g:69,b:45,weight:28},{r:113,g:138,b:189,weight:20},{r:247,g:193,b:132,weight:9},{r:172,g:214,b:141,weight:5},{r:228,g:201,b:218,weight:1}],
  [{r:243,g:208,b:43,weight:593},{r:46,g:151,b:190,weight:552},{r:175,g:54,b:117,weight:348},{r:138,g:195,b:75,weight:222},{r:172,g:208,b:199,weight:169},{r:86,g:99,b:167,weight:129},{r:235,g:102,b:43,weight:105},{r:211,g:135,b:167,weight:70}]]

  process();
};

// mock();
