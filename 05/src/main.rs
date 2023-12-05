use std::str::FromStr;

#[derive(PartialEq)]
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
    destination_start_range: usize,
    source_start_range: usize,
    range: usize
}

impl FromStr for MapData {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut seeds = s.split_ascii_whitespace()
            .map(|data| data.parse::<usize>().expect("Should be a valid number"));

        Ok(MapData{
            destination_start_range: seeds.next().unwrap(),
            source_start_range: seeds.next().unwrap(),
            range: seeds.next().unwrap()
        })
    }
}

struct MapSection {
    kind: MapOf,
    data: Vec<MapData>,
}

impl MapSection {

    fn get_destination(&self, source: usize) -> usize {

        for data in &self.data {
            let destination = data.get_destination_value(source);

            if destination != source {
                return destination;
            }
        }

        source
    }

}

struct Map {
    sections: Vec<MapSection>,
    seeds: Vec<usize>,
}

impl Map {

    fn get_seed_destination(&self, target: MapOf) -> Vec<usize> {
        self.seeds.iter().map(|seed| {
            let mut last_source: usize = *seed;

            for section in &self.sections {
                last_source = section.get_destination(last_source);

                if section.kind == target {
                    return last_source;
                }
            }

            last_source
        })
        .collect::<Vec<_>>()
    }

}

impl FromStr for Map {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut lines = s.lines();
        let mut sections: Vec<MapSection> = vec!();
        let mut current_section: Option<MapSection> = None;

        let seeds = lines.next()
            .expect("Should not be empty")[6..].trim().split_ascii_whitespace()
            .map(|seed| seed.parse::<usize>().expect("Should be a valid number"))
            .collect::<Vec<_>>();

        for line in lines {
            if line.is_empty() {
                if current_section.is_some() {
                    sections.push(current_section.unwrap());
                }

                current_section = None;
            } else if current_section.is_none() {
                let kind = line.parse::<MapOf>().expect("Should be a valid field");

                current_section = Some(MapSection{
                    kind,
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
            seeds,
            sections
        })
    }
}

impl MapData {
    fn get_destination_value(&self, source: usize) -> usize {
        let overlaps = source >= self.source_start_range && source <= self.source_start_range + self.range;

        if !overlaps {
            source
        } else {
            let diff= source - self.source_start_range;
            self.destination_start_range + diff
        }
    }
}


fn main() -> Result<(), std::io::Error> {
    let input = include_str!("input.txt");

    let map = input.parse::<Map>().expect("Should be a valid input");
    let locations = map.get_seed_destination(MapOf::HumidityToLocation);

    println!("{:?}", locations.iter().min());

    Ok(())
}