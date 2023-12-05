use std::ops::Range;
use std::str::FromStr;

enum MapOf {
    SeedToSoil,
    SoilToFertilizer,
    FertilizerToWater,
    WaterToLight,
    LightToTemperature,
    TemperatureToHumidity,
    HumidityToLocation
}

impl FromStr for MapOf {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "seed-to-soil map:"  => Ok(MapOf::SeedToSoil),
            "soil-to-fertilizer map:"  => Ok(MapOf::SoilToFertilizer),
            "fertilizer-to-water map:"  => Ok(MapOf::FertilizerToWater),
            "water-to-light map:"  => Ok(MapOf::WaterToLight),
            "light-to-temperature map:"  => Ok(MapOf::LightToTemperature),
            "temperature-to-humidity map:"  => Ok(MapOf::TemperatureToHumidity),
            "humidity-to-location map:"  => Ok(MapOf::HumidityToLocation),
            _      => Err(()),
        }
    }
}

struct MapData {
    destination_start_range: u128,
    source_start_range: u128,
    range: u128
}

impl FromStr for MapData {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut seeds = s.split_ascii_whitespace()
            .map(|data| data.parse::<u128>().expect("Should be a valid number"));

        Ok(MapData{
            destination_start_range: seeds.next().unwrap(),
            source_start_range: seeds.next().unwrap(),
            range: seeds.next().unwrap()
        })
    }
}

struct MapSection {
    data: Vec<MapData>,
}

impl MapSection {

    fn get_source(&self, destination: u128) -> u128 {

        for data in &self.data {
            let source = data.get_source_value(destination);

            if destination != source {
                return source;
            }
        }

        destination
    }
}

struct Map {
    sections: Vec<MapSection>,
    range_seeds: Vec<Range<u128>>,
}

impl Map {

    fn get_closest_destination(&self) -> u128 {

        for i in 0.. {

            let mut current_destination = i;
            for section in self.sections.iter().rev() {
                current_destination = section.get_source(current_destination);
            }

            if self.range_seeds.iter().any(|range| range.contains(&current_destination)) {
                return i;
            }
        }

        panic!("Should not reach here");
    }

}

impl FromStr for Map {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut lines = s.lines();
        let mut sections: Vec<MapSection> = vec!();
        let mut current_section: Option<MapSection> = None;
        let mut range_seeds: Vec<Range<u128>> = vec!();

        let mut seeds = lines.next()
            .expect("Should not be empty")[6..].trim().split_ascii_whitespace()
            .map(|seed| seed.parse::<u128>().expect("Should be a valid number"));

        while let Some(start_at) = seeds.next() {
            let range = seeds.next().unwrap() - 1;
            range_seeds.push(Range { start: start_at, end: start_at + range });
        }

        for line in lines {
            if line.is_empty() {
                if current_section.is_some() {
                    sections.push(current_section.unwrap());
                }

                current_section = None;
            } else if current_section.is_none() {
                current_section = Some(MapSection{
                    data: vec![],
                });
            } else {
                let data = line.parse::<MapData>().expect("Should be valid");

                current_section.as_mut().unwrap().data.push(data);
            }
        }

        if current_section.is_some() {
            sections.push(current_section.unwrap());
        }

        Ok(Map {
            range_seeds,
            sections
        })
    }
}

impl MapData {
    fn get_source_value(&self, destination: u128) -> u128 {
        let overlaps = destination >= self.destination_start_range && destination <= self.destination_start_range + self.range;

        if !overlaps {
            destination
        } else {
            let diff= destination - self.destination_start_range;
            self.source_start_range + diff
        }
    }
}


fn main() -> Result<(), std::io::Error> {
    let input = include_str!("input.txt");

    let map = input.parse::<Map>().expect("Should be a valid input");
    let closest_location = map.get_closest_destination();

    println!("{:?}", closest_location);

    Ok(())
}