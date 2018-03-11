
```
$ pip install google-cloud-vision
```

```
$ export GOOGLE_APPLICATION_CREDENTIALS="....json"
```

[Cloud Vision API Client Libraries](https://cloud.google.com/vision/docs/reference/libraries?hl=ja#client-libraries-install-python) よりソース引用


```python
import io
import os

# Imports the Google Cloud client library
from google.cloud import vision
from google.cloud.vision import types

# Instantiates a client
client = vision.ImageAnnotatorClient()

# The name of the image file to annotate
file_name = os.path.join(
    os.path.dirname(__file__),
    'resources/wakeupcat.jpg')

# Loads the image into memory
with io.open(file_name, 'rb') as image_file:
    content = image_file.read()

image = types.Image(content=content)

# Performs label detection on the image file
response = client.label_detection(image=image)
labels = response.label_annotations

print('Labels:')
for label in labels:
    print(label.description)
```

```
>>> for label in labels:
...     print(label.description, round(label.score * 100, 1))
...
dog 89.2
dog like mammal 88.7
grass 87.2
dog breed 87.0
grassland 72.5
meadow 65.4
snout 64.8
english cocker spaniel 62.7
spaniel 62.7
pasture 62.5
```

# 参考

[GitHub: Vision - API Reference](https://googlecloudplatform.github.io/google-cloud-python/latest/vision/)

[Google Cloud Vision API: Cloud Vision API Client Libraries](https://cloud.google.com/vision/docs/reference/libraries?hl=ja#client-libraries-install-python)

[Google Cloud Vision API: 顔検出のチュートリアル](https://cloud.google.com/vision/docs/face-tutorial?hl=ja)

[Google Cloud Vision API: Rest Reference](https://cloud.google.com/vision/docs/reference/rest/)

[PyPI:google-cloud-vision](https://pypi.org/project/google-cloud-vision/)
