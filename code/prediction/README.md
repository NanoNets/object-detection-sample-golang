# Code Samples for Object Detection Project

Images and annotations taken from - https://github.com/bourdakos1/Custom-Object-Detection

Images consists of frames taken from a clip from Star Wars: The Force Awakens.
[![Watch the video](https://github.com/bourdakos1/Custom-Object-Detection/raw/master/screenshots/starwars_small.gif)](https://www.youtube.com/watch?v=xW2hpkoaIiM)

Annotations are present for each frame and have the same name as the image name. You can find the example to train a model in golang and node, by updating the api-key and model id in corresponding file. There is also a pre-processed json annotations folder that are ready payload for nanonets api, the script used is node/xml-to-json.js .

How to Train the Millenium Falcon Dataset:
Step 1: Download Code
git clone https://github.com/NanoNets/code-samples.git
cd code-samples

Step 2: Get your API Key
Get your API key here:
http://app.nanonets.com/user/api_key

Step 3 Step API KEY in Env:
Set Environment Variable
export NANONETS_API_KEY=YOUR_API_KEY_GOES_HERE

Step 4 Create a New Model:
go build millenium_falcon/code/create-model && ./create-model
#prints
#now run:
#export NANONETS_MODEL_ID=0000000-0000-0000-00000000

Step 5 Set Model ID in ENV:
export NANONETS_MODEL_ID=YOUR_MODEL_ID

Step 6 Upload Training Data:
go build millenium_falcon/code/upload-training && ./upload-training
#prints responses of uploading multiple images

Step 7 Train the Model:
go build millenium_falcon/code/train-model && ./train-model

Step 8 Get State of Model:
go build millenium_falcon/code/model-state && ./model-state
#prints the state of the model
#wait for the state of the model to become 5
#after this move on to Step 9

Step 9 Make predictions using the Model:
go build millenium_falcon/code/prediction
./prediction PATH_TO_YOUR_IMAGE_TO_PREDICT.jpg

#example usage 
#./prediction ./images/videoplayback0051.jpg