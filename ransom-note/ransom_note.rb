#! /usr/bin/env ruby

# @param {String} ransom_note
# @param {String} magazine
# @return {Boolean}
def can_construct(ransom_note, magazine)
    mag_dict = Hash.new(0)
    magazine.each_char { |c|
        mag_dict[c] += 1
    }
    
    ransom_note.each_char { |c|
        mag_dict[c] -= 1
        if mag_dict[c] < 0 
            return false
        end    
    }
    return true
end

puts can_construct("abc", "abbc")

