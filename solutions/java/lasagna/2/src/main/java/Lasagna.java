public class Lasagna {
    public static final int expectedMinutesInOven(){
        return 40;
    }

    public static final int remainingMinutesInOven(int timeInOven){
        return expectedMinutesInOven() - timeInOven;
    }
    
    public static final int preparationTimeInMinutes(int numberOfLayer){
        return numberOfLayer * 2;
    }

    public static final int totalTimeInMinutes(int numberOfLayer, int timeInOven){
        return preparationTimeInMinutes(numberOfLayer) + timeInOven;
    }
}
