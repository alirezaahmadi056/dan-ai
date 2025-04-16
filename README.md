# 🐦 Lord Shen AI Gateway

<img src=".github/shen.png" alt="Lord Shen" width="400"/>

A simple yet powerful gateway to AI services by barbossa — integrating neural networks, clustering, regression, and LLM models via clean API routes.

---

## 🧠 Models

| Code Name | Model Type         | Description             |
|-----------|--------------------|-------------------------|
| `oogway`  | Neural Network     | For smart, complex tasks |
| `shifu`   | Linear Regression  | For classic predictions |
| `po`      | Clustering         | For unsupervised analysis |

---

## 🚀 Quick Access

### 🔗 Endpoints

| Route | Description |
|-------|-------------|
| [`/`](#) | Home or base route |
| [`/ai/suggestion`](#example-usage-aisuggestion) | Suggestion based on input |
| [`/pridect/comment`](#example-usage-pridectcomment) | Predict feedback sentiment |

> ℹ️ **Note:** All requests must include the `x-api-key` header and a `model` field (`oogway`, `po`, or `shifu`) #(only for /ai/suggestion) when required.

---

## 📬 Headers

Include the following in each request:

```http
x-api-key: YOUR_API_KEY
model: oogway | po | shifu   # (only for /ai/suggestion)
```
## 🛠️ Example

### test

**usage:** <br/>

```bash
    curl https://daneshjooyar.com:12335/
```

### suggestion
    
**usage:** <br/>
send data : 
```json
{
  "age":"18",
  "educationalStatus":"0",//Studying = 0 graduate = 1
  "fieldOfStudy":"1",//computer = 1 more = 0
  "maritalStatus":"0",//married = 1 single = 0
  "gender":"1",//man = 1 woman = 0
  "militaryStatus":"1",//Educational exemption = 1  exemption = 2 runaway = 3 
  "freeTime":"2",//with hours
  "targetIncome":"14",//with milion
  "intentionToMigrate":"0",//true = 1 false = 0
  "interestInMathematics":"0",//true = 1 false = 0
  "computerExperience":"1",//with year
  "whichOneDoYouLikeMore":"0",//engineering = 1 artistic = 0
  "whichCaseIsmoreRelevant":"0",//fast = 0 slow = 1
  "doYouWorkOnHolidays":"1",//true = 1 false = 0
  "disability":"0",//true = 1 false = 0
  "addictionred":"0"//true = 1 false = 0
}
```
result data:
```json
{
  "result": [
    "android",
    "ai",
    "wordpress"
  ]
}
```

### predict comment

**usage:** <br/>
send data :
```json
{
  "comments": [
    "comment0",
    "comment1",
    "and so on"
  ] 
}
```
result data:
```json
{
  "result":"it is a good course" 
}
```

# ⚠️ Errors

## Error Codes
**404:** <br/>
not found server
<br/>

**401:** <br/>
invalid api key
<br/>

**400:** <br/>
bad request
<br/>


**500:** <br/>
internal server error `#(like crash models or crash ollama and so on)`
<br/>


**503:** <br/>
service unavailable
<br/>


    
    
