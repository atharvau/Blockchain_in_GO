const backgroundColor = 0x000000
// $('#athrva').animate({ left: '130px' })

//methods
/*////////////////////////////////////////*/
let clock = new THREE.Clock()
var renderCalls = []
function render() {
  requestAnimationFrame(render)
  renderCalls.forEach(callback => {
    callback()
  })
}
render()

/*////////////////////////////////////////*/

var scene = new THREE.Scene()
let mixer = null
var camera = new THREE.PerspectiveCamera(2, window.innerWidth / window.innerHeight, 0.1, 800)
camera.position.set(-300, -300, -300)

var renderer = new THREE.WebGLRenderer({ alpha: true })
renderer.setPixelRatio(window.devicePixelRatio)
renderer.setSize(window.innerWidth, window.innerHeight)
renderer.setClearColor(0x000000, 0) // the default
renderer.sortObjects = false
renderer.toneMapping = THREE.LinearToneMapping
renderer.toneMappingExposure = Math.pow(0.94, 5.0)
renderer.shadowMap.enabled = true
renderer.shadowMap.type = THREE.PCFShadowMap

var container = document.getElementById('canvas')

window.addEventListener(
  'resize',
  function() {
    camera.aspect = window.innerWidth / window.innerHeight
    camera.updateProjectionMatrix()
    renderer.setSize(window.innerWidth, window.innerHeight)
  },
  false
)

container.appendChild(renderer.domElement)

function renderScene() {
  renderer.render(scene, camera)
}
renderCalls.push(renderScene)

/* ////////////////////////////////////////////////////////////////////////// */

var controls = new THREE.OrbitControls(camera)

controls.rotateSpeed = 0.3
controls.zoomSpeed = 0.9

controls.minDistance = 30
controls.maxDistance = 20

controls.minPolarAngle = 0 // radians
controls.maxPolarAngle = Math.PI / 2 // radians

controls.enableDamping = false
controls.enablePan = false

// for (var i = 0; i < gltf.animations.length; i++) {
controls.dampingFactor = 0.05
// if youre definition is like
// you can easily disable it by using
renderCalls.push(function() {
  controls.update()
})

var light = new THREE.AmbientLight(0xffffff) // soft white light
scene.add(light)

var directionalLight2 = new THREE.DirectionalLight(0xffffff, 2)
directionalLight2.position.set(-4, 0, 0)

var directionalLight = new THREE.DirectionalLight(0xffffff, 2)
directionalLight.position.set(5, 0, 0)
// scene.add(directionalLight2)
scene.add(directionalLight)
/* ////////////////////////////////////////////////////////////////////////// */

/* ////////////////////////////////////////////////////////////////////////// */
var object
var gltf
var loader = new THREE.GLTFLoader()
loader.crossOrigin = true
loader.load('./templates/scene.gltf', function(data) {
  bject = data.scene
  object.depthWrite = false

  object.scale.set(5, 7, 8);
  object.position.set(-0.5, -0.7, 0)
  object.rotation.set(-0.3, 2.4, 0.4)
  renderer.render(scene, camera)


  object.position.set(-0.1, 0, 0)
  object.rotation.set(-0.3, 2.9, 0.4)


  object.position.y = -0.2

  gltf = data
  if (gltf.animations && gltf.animations.length) {
    mixer = new THREE.AnimationMixer(gltf.scene)
    var animation = gltf.animations[0]
    mixer.clipAction(animation).play()
  }

  // }

  scene.add(object)
  //, onProgress, onError );
})

function animate() {
  requestAnimationFrame(animate)
   object.rotation.y += 0.002
  controls.update()
  if (mixer) {
    mixer.update(clock.getDelta() * mixer.timeScale)
  }
  renderer.render(scene, camera)
  camera.lookAt(scene)
}
animate()

count = 0
var boola2 = true
var jamna = 0
function hu(e, delta) {
  // object.rotation.set(0, 1, 0)
}
