class Student extends Person {

    private int[] testScores;

    Student(String firstName, String lastName, int idNumber, int[] testScores) {
        super(firstName, lastName, idNumber);
        this.testScores = testScores;
    }

    public void printStudent() {

    }

    public char calculate() {
        char result = 0;
        int average;
        int sum = 0;
        for (int i = 0; i < testScores.length; i++) {
            sum = testScores[i] + sum;
        }

        average = sum / (testScores.length);
        if (average < 40) {
            result = 'T';
        } else if (average >= 40 && average < 55) {
            result = 'D';
        } else if (average >= 70 && average < 80) {
            result = 'A';
        } else if (average >= 55 && average < 70) {
            result = 'P';
        } else if (average >= 80 && average < 90) {
            result = 'E';
        } else if (average >= 90 && average <= 100) {
            result = 'O';
        }

        return result;

    }

}