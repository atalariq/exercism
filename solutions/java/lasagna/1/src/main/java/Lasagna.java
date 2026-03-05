public class Lasagna {
    // TODO: define the 'expectedMinutesInOven()' method
    public static final int expectedMinutesInOven(){
        return 40;
    }

    // TODO: define the 'remainingMinutesInOven()' method
    public static final int remainingMinutesInOven(int timeInOven){
        return expectedMinutesInOven() - timeInOven;
    }
    

    // TODO: define the 'preparationTimeInMinutes()' method
    public static final int preparationTimeInMinutes(int numberOfLayer){
        return numberOfLayer * 2;
    }

    // TODO: define the 'totalTimeInMinutes()' method
    public static final int totalTimeInMinutes(int numberOfLayer, int timeInOven){
        return preparationTimeInMinutes(numberOfLayer) + timeInOven;
    }
}
