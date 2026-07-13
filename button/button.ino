const int BUTTON_PIN = 2;
const int GREEN_LED = 9;
const int RED_LED = 11;
const int YELLOW_LED = 4;

int lastState = HIGH;
int confirmedState = HIGH;
unsigned long lastDebounceTime = 0;
const unsigned long DEBOUNCE_DELAY = 50;

bool greenOn = false;
bool redOn = false;
bool yellowOn = false;
unsigned long greenOffTime = 0;
unsigned long redOffTime = 0;
unsigned long yellowOffTime = 0;


void setup() {
  pinMode(BUTTON_PIN, INPUT_PULLUP);
  pinMode(GREEN_LED, OUTPUT);
  pinMode(RED_LED, OUTPUT);
  pinMode(YELLOW_LED, OUTPUT);
  Serial.begin(9600);
}

void loop() {
  checkButton();
  checkFeedback();
  updateLeds();
}


void checkButton() {
  int currentState = digitalRead(BUTTON_PIN);
  if (currentState != lastState) {
    lastDebounceTime = millis();
  }
  if (millis() - lastDebounceTime > DEBOUNCE_DELAY) {
    if (currentState != confirmedState) {
      confirmedState = currentState;
      if (confirmedState == LOW) {
        Serial.println("ALARM");
      }
    }
  }
  lastState = currentState;
}


void checkFeedback() {
  if (Serial.available() > 0) {
    String response = Serial.readStringUntil('\n');
    if (response == "OK") {
      digitalWrite(GREEN_LED, HIGH);
      greenOffTime = millis() + 1000;
      greenOn = true;
    } else if (response == "FAIL") {
       for (int i = 0; i < 3; i++) {
        digitalWrite(RED_LED, HIGH);
        delay(100);
        digitalWrite(RED_LED, LOW);
        delay(100);
      } 
    } else {
        digitalWrite(YELLOW_LED, HIGH);
        yellowOffTime = millis() + 1000;
        yellowOn = true;
      }
    }
}

void updateLeds() {
  if (greenOn && millis() >= greenOffTime) {
    digitalWrite(GREEN_LED, LOW);
    greenOn = false;
  }
  if (redOn && millis() >= redOffTime) {
    redOn = false;
  }
  if (yellowOn && millis() >= yellowOffTime) {
    digitalWrite(YELLOW_LED, LOW);
    yellowOn = false;
  }
}


