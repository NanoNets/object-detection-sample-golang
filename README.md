# Code Samples for Object Detection Project Written in Golang

Images and annotations taken from - https://github.com/bourdakos1/Custom-Object-Detection

Images consists of frames taken from a clip from Star Wars: The Force Awakens.
[![Watch the video](https://github.com/bourdakos1/Custom-Object-Detection/raw/master/screenshots/starwars_small.gif)](https://www.youtube.com/watch?v=xW2hpkoaIiM)

Annotations are present for each frame and have the same name as the image name. You can find the example to train a model in golang and node, by updating the api-key and model id in corresponding file. There is also a pre-processed json annotations folder that are ready payload for nanonets api, the script used is node/xml-to-json.js .

 >**Note:** Make sure you have go installed on your system if you don't. Install it from: https://golang.org/doc/install
 
 
### Step 1: Clone the Repo
```bash
cd $GOPATH/src
git clone https://github.com/NanoNets/object-detection-sample-golang.git
cd object-detection-sample-golang
```

### Step 2: Get your API Key
Get your API Key from http://app.nanonets.com/user/api_key

### Step 3: Set the API key as an Environment Variable
```bash
export NANONETS_API_KEY=YOUR_API_KEY_GOES_HERE
```

### Step 4: Create a New Model
```bash
go build object-detection-sample-golang/code/create-model && ./create-model
```
 >_**Note:** This generates a MODEL_ID that you need for the next step

### Step 5: Add Model Id as Environment Variable
```bash
export NANONETS_MODEL_ID=YOUR_MODEL_ID
```
 >_**Note:** you will get YOUR_MODEL_ID from the previous step

### Step 6: Upload the Training Data
The training data is found in images (image files) and annotations (annotations for the files)
```bash
go build object-detection-sample-golang/code/upload-training && ./upload-training
```

### Step 7: Train Model
Once the Images have been uploaded, begin training the Model
```bash
go build object-detection-sample-golang/code/train-model && ./train-model
```

### Step 8: Get Model State
The model takes approximately an hour to train. You will get an email once the model is trained. In the meanwhile you check the state of the model
```bash
go build object-detection-sample-golang/code/model-state && ./model-state
```

### Step 9: Make Prediction
Once the model is trained. You can make predictions using the model
```bash
go build object-detection-sample-golang/code/prediction
./prediction PATH_TO_YOUR_IMAGE.jpg
```

**Sample Usage:**
```bash
./prediction ./images/videoplayback0051.jpg
```


*Note the golang sample uses the comverted json instead of the xml payload for convenience purposes, hence it has no dependencies.*
